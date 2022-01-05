%{
package query

import (
  "regexp"
)

var lastQuery Query

%}

%union {
  qry Query
  fix int64
  flo float64
  str string
  rxp *regexp.Regexp
}

%type <qry> query query1 query2 top keyval
%type <fix> fix
%type <flo> flo
%token '=' '>' '<' '(' ')' '!' '&' '|' '-'
%token LE GE NE
%token <fix> FIX
%token <flo> FLO
%token <str> WORD
%token <str> STR
%token <rxp> RXP

%%

top:
  query {
    // $$ = $1
    lastQuery = $1
  }

query:
	query1
|	'!' query
	{
		$$ = NewNegQuery($2)
	}

query1:
	query2
|	query1 '&' query2
	{
		$$ = NewAndQuery($1, $3)
	}
|	query1 '|' query2
	{
		$$ = NewOrQuery($1, $3)
	}

query2:
       RXP {
              $$ = NewRegexpQuery($1)
       }
|      STR {
              $$ = NewInQuery($1)
       }
|      WORD {
              $$ = NewInQuery($1)
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
              $$ = NewKeyQuery(Eq, $1, $3)
       }
|      WORD '=' WORD {
              $$ = NewKeyQuery(Eq, $1, $3)
       }
|      WORD '=' RXP {
              $$ = NewKeyQuery(Eq, $1, $3)
       }
|      WORD '=' fix {
              $$ = NewKeyQuery(Eq, $1, $3)
       }
|      WORD '=' flo {
              $$ = NewKeyQuery(Eq, $1, $3)
       }
|       WORD NE STR {
              $$ = NewKeyQuery(Ne, $1, $3)
       }
|      WORD NE WORD {
              $$ = NewKeyQuery(Ne, $1, $3)
       }
|      WORD NE RXP {
              $$ = NewKeyQuery(Ne, $1, $3)
       }
|      WORD NE fix {
              $$ = NewKeyQuery(Ne, $1, $3)
       }
|      WORD NE flo {
              $$ = NewKeyQuery(Ne, $1, $3)
       }
|      WORD '>' fix {
              $$ = NewKeyQuery(Gt, $1, $3)
       }
|      WORD '>' flo {
              $$ = NewKeyQuery(Gt, $1, $3)
       }
|      WORD '<' fix {
              $$ = NewKeyQuery(Lt, $1, $3)
       }
|      WORD '<' flo {
              $$ = NewKeyQuery(Lt, $1, $3)
       }
|      WORD GE fix {
              $$ = NewKeyQuery(Ge, $1, $3)
       }
|      WORD GE flo {
              $$ = NewKeyQuery(Ge, $1, $3)
       }
|      WORD LE fix {
              $$ = NewKeyQuery(Le, $1, $3)
       }
|      WORD LE flo {
              $$ = NewKeyQuery(Le, $1, $3)
       }


fix:
      '-' FIX { $$ = -1 * $2 }
|     FIX { $$ = $1 }

flo:
'-' FLO { $$ = float64(-1.0) * $2 }
|     FLO { $$ = $1 }

%%
