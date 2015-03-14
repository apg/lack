%{
package query

import (
  "regexp"
)

var lastMatcher Matcher

%}

%union {
  mat Matcher
  fix int64
  flo float64
  str string
  rxp *regexp.Regexp
}

%type <mat> query query1 query2 top keyval
%type <fix> fix
%type <flo> flo
%token '=' '>' '<' '(' ')' '!' '&' '|' '-'
%token LE GE NE
%token <fix> FIX
nn%token <flo> FLO
%token <str> WORD
%token <str> STR
%token <rxp> RXP

%%

top:
  query {
    // $$ = $1
    lastMatcher = $1
  }

query:
	query1
|	'!' query
	{
		$$ = NewNegMatcher($2)
	}

query1:
	query2
|	query1 '&' query2
	{
		$$ = NewAndMatcher($1, $3)
	}
|	query1 '|' query2
	{
		$$ = NewOrMatcher($1, $3)
	}

query2:
       RXP {
              $$ = NewRegexpMatcher($1)
       }
|      STR {
              $$ = NewInMatcher($1)
       }
|      WORD {
              $$ = NewInMatcher($1)
       }
|       keyval {
              $$ = $1
       }
|	'(' query ')'
	{
		$$ = $2
	}

keyval:
       WORD '=' STR {
              $$ = NewKeyMatcher(Eq, $1, $3)
       }
|      WORD '=' WORD {
              $$ = NewKeyMatcher(Eq, $1, $3)
       }
|      WORD '=' RXP {
              $$ = NewKeyMatcher(Eq, $1, $3)
       }
|      WORD '=' fix {
              $$ = NewKeyMatcher(Eq, $1, $3)
       }
|      WORD '=' flo {
              $$ = NewKeyMatcher(Eq, $1, $3)
       }
|       WORD NE STR {
              $$ = NewKeyMatcher(Ne, $1, $3)
       }
|      WORD NE WORD {
              $$ = NewKeyMatcher(Ne, $1, $3)
       }
|      WORD NE RXP {
              $$ = NewKeyMatcher(Ne, $1, $3)
       }
|      WORD NE fix {
              $$ = NewKeyMatcher(Ne, $1, $3)
       }
|      WORD NE flo {
              $$ = NewKeyMatcher(Ne, $1, $3)
       }
|      WORD '>' fix {
              $$ = NewKeyMatcher(Gt, $1, $3)
       }
|      WORD '>' flo {
              $$ = NewKeyMatcher(Gt, $1, $3)
       }
|      WORD '<' fix {
              $$ = NewKeyMatcher(Lt, $1, $3)
       }
|      WORD '<' flo {
              $$ = NewKeyMatcher(Lt, $1, $3)
       }
|      WORD GE fix {
              $$ = NewKeyMatcher(Ge, $1, $3)
       }
|      WORD GE flo {
              $$ = NewKeyMatcher(Ge, $1, $3)
       }
|      WORD LE fix {
              $$ = NewKeyMatcher(Le, $1, $3)
       }
|      WORD LE flo {
              $$ = NewKeyMatcher(Le, $1, $3)
       }


fix:
      '-' FIX { $$ = -1 * $2 }
|     FIX { $$ = $1 }

flo:
'-' FLO { $$ = float64(-1.0) * $2 }
|     FLO { $$ = $1 }

%%
