//line query.y:2
package query

import __yyfmt__ "fmt"

//line query.y:2
import (
	"regexp"
)

var lastMatcher Matcher

//line query.y:12
type querySymType struct {
	yys int
	mat Matcher
	fix int64
	flo float64
	str string
	rxp *regexp.Regexp
}

const LE = 57346
const GE = 57347
const NE = 57348
const FIX = 57349
const nn = 57350
const FLO = 57351
const WORD = 57352
const STR = 57353
const RXP = 57354

var queryToknames = []string{
	"'='",
	"'>'",
	"'<'",
	"'('",
	"')'",
	"'!'",
	"'&'",
	"'|'",
	"'-'",
	"LE",
	"GE",
	"NE",
	"FIX",
	"nn",
	"FLO",
	"WORD",
	"STR",
	"RXP",
}
var queryStatenames = []string{}

const queryEofCode = 1
const queryErrCode = 2
const queryMaxDepth = 200

//line query.y:140

//line yacctab:1
var queryExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const queryNprod = 34
const queryPrivate = 57344

var queryTokenNames []string
var queryStates []string

const queryLast = 67

var queryAct = []int{

	28, 45, 44, 46, 29, 9, 30, 32, 31, 33,
	28, 10, 1, 4, 29, 10, 30, 24, 23, 25,
	27, 5, 3, 8, 7, 6, 26, 8, 7, 6,
	11, 12, 0, 21, 22, 0, 35, 37, 39, 41,
	43, 0, 34, 36, 38, 40, 42, 28, 14, 16,
	17, 29, 0, 30, 0, 2, 0, 19, 18, 15,
	13, 0, 0, 0, 0, 0, 20,
}
var queryPact = []int{

	4, -1000, -1000, 20, 4, -1000, -1000, -1000, 44, -1000,
	4, 8, 8, -1000, -2, -12, 35, 35, 35, 35,
	-6, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -15, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var queryPgo = []int{

	0, 55, 22, 21, 12, 5, 26, 20,
}
var queryR1 = []int{

	0, 4, 1, 1, 2, 2, 2, 3, 3, 3,
	3, 3, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	6, 6, 7, 7,
}
var queryR2 = []int{

	0, 1, 1, 2, 1, 3, 3, 1, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 1, 2, 1,
}
var queryChk = []int{

	-1000, -4, -1, -2, 9, -3, 21, 20, 19, -5,
	7, 10, 11, -1, 4, 15, 5, 6, 14, 13,
	-1, -3, -3, 20, 19, 21, -6, -7, 12, 16,
	18, 20, 19, 21, -6, -7, -6, -7, -6, -7,
	-6, -7, -6, -7, 8, 16, 18,
}
var queryDef = []int{

	0, -2, 1, 2, 0, 4, 7, 8, 9, 10,
	0, 0, 0, 3, 0, 0, 0, 0, 0, 0,
	0, 5, 6, 12, 13, 14, 15, 16, 0, 31,
	33, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 11, 30, 32,
}
var queryTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 9, 3, 3, 3, 3, 10, 3,
	7, 8, 3, 3, 3, 12, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	6, 4, 5, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 11,
}
var queryTok2 = []int{

	2, 3, 13, 14, 15, 16, 17, 18, 19, 20,
	21,
}
var queryTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var queryDebug = 0

type queryLexer interface {
	Lex(lval *querySymType) int
	Error(s string)
}

const queryFlag = -1000

func queryTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(queryToknames) {
		if queryToknames[c-4] != "" {
			return queryToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func queryStatname(s int) string {
	if s >= 0 && s < len(queryStatenames) {
		if queryStatenames[s] != "" {
			return queryStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func querylex1(lex queryLexer, lval *querySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = queryTok1[0]
		goto out
	}
	if char < len(queryTok1) {
		c = queryTok1[char]
		goto out
	}
	if char >= queryPrivate {
		if char < queryPrivate+len(queryTok2) {
			c = queryTok2[char-queryPrivate]
			goto out
		}
	}
	for i := 0; i < len(queryTok3); i += 2 {
		c = queryTok3[i+0]
		if c == char {
			c = queryTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = queryTok2[1] /* unknown char */
	}
	if queryDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", queryTokname(c), uint(char))
	}
	return c
}

func queryParse(querylex queryLexer) int {
	var queryn int
	var querylval querySymType
	var queryVAL querySymType
	queryS := make([]querySymType, queryMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	querystate := 0
	querychar := -1
	queryp := -1
	goto querystack

ret0:
	return 0

ret1:
	return 1

querystack:
	/* put a state and value onto the stack */
	if queryDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", queryTokname(querychar), queryStatname(querystate))
	}

	queryp++
	if queryp >= len(queryS) {
		nyys := make([]querySymType, len(queryS)*2)
		copy(nyys, queryS)
		queryS = nyys
	}
	queryS[queryp] = queryVAL
	queryS[queryp].yys = querystate

querynewstate:
	queryn = queryPact[querystate]
	if queryn <= queryFlag {
		goto querydefault /* simple state */
	}
	if querychar < 0 {
		querychar = querylex1(querylex, &querylval)
	}
	queryn += querychar
	if queryn < 0 || queryn >= queryLast {
		goto querydefault
	}
	queryn = queryAct[queryn]
	if queryChk[queryn] == querychar { /* valid shift */
		querychar = -1
		queryVAL = querylval
		querystate = queryn
		if Errflag > 0 {
			Errflag--
		}
		goto querystack
	}

querydefault:
	/* default state action */
	queryn = queryDef[querystate]
	if queryn == -2 {
		if querychar < 0 {
			querychar = querylex1(querylex, &querylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if queryExca[xi+0] == -1 && queryExca[xi+1] == querystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			queryn = queryExca[xi+0]
			if queryn < 0 || queryn == querychar {
				break
			}
		}
		queryn = queryExca[xi+1]
		if queryn < 0 {
			goto ret0
		}
	}
	if queryn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			querylex.Error("syntax error")
			Nerrs++
			if queryDebug >= 1 {
				__yyfmt__.Printf("%s", queryStatname(querystate))
				__yyfmt__.Printf(" saw %s\n", queryTokname(querychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for queryp >= 0 {
				queryn = queryPact[queryS[queryp].yys] + queryErrCode
				if queryn >= 0 && queryn < queryLast {
					querystate = queryAct[queryn] /* simulate a shift of "error" */
					if queryChk[querystate] == queryErrCode {
						goto querystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if queryDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", queryS[queryp].yys)
				}
				queryp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if queryDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", queryTokname(querychar))
			}
			if querychar == queryEofCode {
				goto ret1
			}
			querychar = -1
			goto querynewstate /* try again in the same state */
		}
	}

	/* reduction by production queryn */
	if queryDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", queryn, queryStatname(querystate))
	}

	querynt := queryn
	querypt := queryp
	_ = querypt // guard against "declared and not used"

	queryp -= queryR2[queryn]
	queryVAL = queryS[queryp+1]

	/* consult goto table to find next state */
	queryn = queryR1[queryn]
	queryg := queryPgo[queryn]
	queryj := queryg + queryS[queryp].yys + 1

	if queryj >= queryLast {
		querystate = queryAct[queryg]
	} else {
		querystate = queryAct[queryj]
		if queryChk[querystate] != -queryn {
			querystate = queryAct[queryg]
		}
	}
	// dummy call; replaced with literal code
	switch querynt {

	case 1:
		//line query.y:34
		{
			// $$ = $1
			lastMatcher = queryS[querypt-0].mat
		}
	case 2:
		queryVAL.mat = queryS[querypt-0].mat
	case 3:
		//line query.y:42
		{
			queryVAL.mat = NewNegMatcher(queryS[querypt-0].mat)
		}
	case 4:
		queryVAL.mat = queryS[querypt-0].mat
	case 5:
		//line query.y:49
		{
			queryVAL.mat = NewAndMatcher(queryS[querypt-2].mat, queryS[querypt-0].mat)
		}
	case 6:
		//line query.y:53
		{
			queryVAL.mat = NewOrMatcher(queryS[querypt-2].mat, queryS[querypt-0].mat)
		}
	case 7:
		//line query.y:58
		{
			queryVAL.mat = NewRegexpMatcher(queryS[querypt-0].rxp)
		}
	case 8:
		//line query.y:61
		{
			queryVAL.mat = NewInMatcher(queryS[querypt-0].str)
		}
	case 9:
		//line query.y:64
		{
			queryVAL.mat = NewInMatcher(queryS[querypt-0].str)
		}
	case 10:
		//line query.y:67
		{
			queryVAL.mat = queryS[querypt-0].mat
		}
	case 11:
		//line query.y:71
		{
			queryVAL.mat = queryS[querypt-1].mat
		}
	case 12:
		//line query.y:76
		{
			queryVAL.mat = NewKeyMatcher(Eq, queryS[querypt-2].str, queryS[querypt-0].str)
		}
	case 13:
		//line query.y:79
		{
			queryVAL.mat = NewKeyMatcher(Eq, queryS[querypt-2].str, queryS[querypt-0].str)
		}
	case 14:
		//line query.y:82
		{
			queryVAL.mat = NewKeyMatcher(Eq, queryS[querypt-2].str, queryS[querypt-0].rxp)
		}
	case 15:
		//line query.y:85
		{
			queryVAL.mat = NewKeyMatcher(Eq, queryS[querypt-2].str, queryS[querypt-0].fix)
		}
	case 16:
		//line query.y:88
		{
			queryVAL.mat = NewKeyMatcher(Eq, queryS[querypt-2].str, queryS[querypt-0].flo)
		}
	case 17:
		//line query.y:91
		{
			queryVAL.mat = NewKeyMatcher(Ne, queryS[querypt-2].str, queryS[querypt-0].str)
		}
	case 18:
		//line query.y:94
		{
			queryVAL.mat = NewKeyMatcher(Ne, queryS[querypt-2].str, queryS[querypt-0].str)
		}
	case 19:
		//line query.y:97
		{
			queryVAL.mat = NewKeyMatcher(Ne, queryS[querypt-2].str, queryS[querypt-0].rxp)
		}
	case 20:
		//line query.y:100
		{
			queryVAL.mat = NewKeyMatcher(Ne, queryS[querypt-2].str, queryS[querypt-0].fix)
		}
	case 21:
		//line query.y:103
		{
			queryVAL.mat = NewKeyMatcher(Ne, queryS[querypt-2].str, queryS[querypt-0].flo)
		}
	case 22:
		//line query.y:106
		{
			queryVAL.mat = NewKeyMatcher(Gt, queryS[querypt-2].str, queryS[querypt-0].fix)
		}
	case 23:
		//line query.y:109
		{
			queryVAL.mat = NewKeyMatcher(Gt, queryS[querypt-2].str, queryS[querypt-0].flo)
		}
	case 24:
		//line query.y:112
		{
			queryVAL.mat = NewKeyMatcher(Lt, queryS[querypt-2].str, queryS[querypt-0].fix)
		}
	case 25:
		//line query.y:115
		{
			queryVAL.mat = NewKeyMatcher(Lt, queryS[querypt-2].str, queryS[querypt-0].flo)
		}
	case 26:
		//line query.y:118
		{
			queryVAL.mat = NewKeyMatcher(Ge, queryS[querypt-2].str, queryS[querypt-0].fix)
		}
	case 27:
		//line query.y:121
		{
			queryVAL.mat = NewKeyMatcher(Ge, queryS[querypt-2].str, queryS[querypt-0].flo)
		}
	case 28:
		//line query.y:124
		{
			queryVAL.mat = NewKeyMatcher(Le, queryS[querypt-2].str, queryS[querypt-0].fix)
		}
	case 29:
		//line query.y:127
		{
			queryVAL.mat = NewKeyMatcher(Le, queryS[querypt-2].str, queryS[querypt-0].flo)
		}
	case 30:
		//line query.y:133
		{
			queryVAL.fix = -1 * queryS[querypt-0].fix
		}
	case 31:
		//line query.y:134
		{
			queryVAL.fix = queryS[querypt-0].fix
		}
	case 32:
		//line query.y:137
		{
			queryVAL.flo = float64(-1.0) * queryS[querypt-0].flo
		}
	case 33:
		//line query.y:138
		{
			queryVAL.flo = queryS[querypt-0].flo
		}
	}
	goto querystack /* stack new state and value */
}
