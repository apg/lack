// Copyright (C) 2013 Keith Rarick, Blake Mizerany

// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use, copy,
// modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
// BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
// ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package logfmt

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var timeFormats = []string{
	time.RFC3339,
	time.RFC3339Nano,
	time.RFC1123,
	time.RFC1123Z,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}

// ErrUnterminatedString is returned when improperly quoted strings are discovered.
var ErrUnterminatedString = errors.New("logfmt: unterminated string")

// Scan returns a log line, parsed as logfmt.
func Scan(line *Line) (err error) {
	var c byte
	var i int
	var m int
	var key []byte
	var val []byte
	var ok bool
	var esc bool

	data := line.line

garbage:
	if i == len(data) {
		return
	}

	c = data[i]
	switch {
	case c > ' ' && c != '"' && c != '=':
		key, val = nil, nil
		m = i
		i++
		goto key
	default:
		i++
		goto garbage
	}

key:
	if i >= len(data) {
		if m >= 0 {
			key = data[m:i]
			line.Append(string(key), nil)
		}
		return
	}

	c = data[i]
	switch {
	case c > ' ' && c != '"' && c != '=':
		i++
		goto key
	case c == '=':
		key = data[m:i]
		i++
		goto equal
	default:
		key = data[m:i]
		i++
		line.Append(string(key), nil)
		goto garbage
	}

equal:
	if i >= len(data) {
		if m >= 0 {
			i--
			key = data[m:i]
			line.Append(string(key), nil)
		}
		return
	}

	c = data[i]
	switch {
	case c > ' ' && c != '"' && c != '=':
		m = i
		i++
		goto ivalue
	case c == '"':
		m = i
		i++
		esc = false
		goto qvalue
	default:
		if key != nil {
			line.Append(string(key), val)
		}
		i++
		goto garbage
	}

ivalue:
	if i >= len(data) {
		if m >= 0 {
			val = data[m:i]
			line.Append(string(key), val)
		}
		return
	}

	c = data[i]
	switch {
	case c > ' ' && c != '"' && c != '=':
		i++
		goto ivalue
	default:
		val = data[m:i]
		line.Append(string(key), val)
		i++
		goto garbage
	}

qvalue:
	if i >= len(data) {
		if m >= 0 {
			err = ErrUnterminatedString
		}
		return
	}

	c = data[i]
	switch c {
	case '\\':
		i += 2
		esc = true
		goto qvalue
	case '"':
		i++
		val = data[m:i]
		if esc {
			val, ok = unquoteBytes(val)
			if !ok {
				err = fmt.Errorf("logfmt: error unquoting bytes %q", string(val))
				goto garbage
			}
		} else {
			val = val[1 : len(val)-1]
		}
		line.Append(string(key), val)
		goto garbage
	default:
		i++
		goto qvalue
	}
}

func parseValue(val []byte) interface{} {
	s := string(val)
	if fix, err := strconv.ParseInt(s, 10, 64); err == nil {
		return fix
	} else if flo, err := strconv.ParseFloat(s, 64); err == nil {
		return flo
		//} else if tim, err := tryTimes(s); err == nil {
		//	return tim
	}
	return s
}

func tryTimes(s string) (time.Time, error) {
	var err error
	var t time.Time
	for _, f := range timeFormats {
		if t, err = time.Parse(f, s); err == nil {
			return t, nil
		}
	}
	return t, err
}
