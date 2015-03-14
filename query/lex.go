package query

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
	"unicode/utf8"
)

type queryLex struct {
	line []byte
	peek rune
}

func (x *queryLex) Lex(yylval *querySymType) int {
	for {
		c := x.next()
		switch c {
		case eof:
			return eof
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return x.num(c, yylval)
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
			'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
			'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			return x.word(c, yylval)
		case '=', '(', ')', '&', '|', '-':
			return int(c)
		case '"':
			return x.str(c, yylval)
		case '>':
			d := x.next()
			if d == '=' {
				return GE
			}
			x.peek = d
			return int(c)
		case '<':
			d := x.next()
			if d == '=' {
				return LE
			}
			x.peek = d
			return int(c)
		case '!':
			d := x.next()
			if d == '=' {
				return NE
			}
			x.peek = d
			return int(c)
		case '/':
			return x.rxp(c, yylval)
		case ' ':
			continue
		default:
			x.Error("unknown character!")
		}
	}
}

func (x *queryLex) Error(s string) {
	log.Fatalf("parse error: %s", s)
}

func (x *queryLex) num(c rune, yylval *querySymType) int {
	var b bytes.Buffer
	addRune(&b, c)
L:
	for {
		c = x.next()
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E':
			addRune(&b, c)
		default:
			break L
		}
	}
	if c != eof {
		x.peek = c
	}

	if fix, err := strconv.ParseInt(b.String(), 10, 64); err == nil {
		yylval.fix = fix
		return FIX
	} else if flo, err := strconv.ParseFloat(b.String(), 64); err == nil {
		yylval.flo = flo
		return FLO
	} else {
		log.Printf("bad number %q", b.String())
		return eof
	}
}

func (x *queryLex) word(c rune, yylval *querySymType) int {
	var b bytes.Buffer
	addRune(&b, c)
L:
	for {
		c = x.next()
		switch {
		case (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'):
			addRune(&b, c)
		case c == '#' || c == '.': // specials.
			addRune(&b, c)
		default:
			break L
		}
	}
	if c != eof {
		x.peek = c
	}

	yylval.str = b.String()
	return WORD
}

func (x *queryLex) str(c rune, yylval *querySymType) int {
	var b bytes.Buffer
L:
	for {
		c = x.next()
		switch c {
		case eof:
			log.Fatalf("Unclosed string.")
		case '\\':
			d := x.next()
			switch d {
			case '\\':
				addRune(&b, d)
				x.peek = eof
			case '"':
				addRune(&b, d)
				x.peek = eof
			default:
				log.Fatalf("Improper escape sequence.")
			}
		case '"':
			c = eof
			break L
		default:
			addRune(&b, c)
		}
	}

	if c != eof {
		x.peek = c
	}

	yylval.str = b.String()
	return STR
}

func (x *queryLex) rxp(c rune, yylval *querySymType) int {
	var b bytes.Buffer
L:
	for {
		c = x.next()
		switch c {
		case eof:
			log.Fatalf("Unclosed regexp.")
		case '\\':
			d := x.next()
			switch d {
			case '/':
				addRune(&b, d)
				x.peek = eof
			default:
				log.Fatalf("Improper escape sequence.")
			}
		case '/':
			c = eof
			break L
		default:
			addRune(&b, c)
		}
	}

	if c != eof {
		x.peek = c
	}

	re, err := regexp.CompilePOSIX(b.String())
	if err != nil {
		log.Fatalf("Failed to compile regular expression: %s\n", err)
	}

	yylval.rxp = re
	return RXP
}

// Return the next rune for the lexer.
func (x *queryLex) next() rune {
	if x.peek != eof {
		r := x.peek
		x.peek = eof
		return r
	}
	if len(x.line) == 0 {
		return eof
	}
	c, size := utf8.DecodeRune(x.line)
	x.line = x.line[size:]
	if c == utf8.RuneError && size == 1 {
		log.Print("invalid utf8")
		return x.next()
	}
	return c
}

func addRune(b *bytes.Buffer, c rune) {
	if _, err := b.WriteRune(c); err != nil {
		log.Fatalf("WriteRune: %s", err)
	}
}
