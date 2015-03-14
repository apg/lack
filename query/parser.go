//go:generate -command yacc go tool yacc
//go:generate yacc -o query.go -p "query" query.y

package query

import "errors"

const eof = 0

// Parse returns a parsed query
func Parse(q []byte) (Matcher, error) {
	if queryParse(&queryLex{line: q}) == 0 {
		return lastMatcher, nil
	}
	return nil, errors.New("Parse error!")
}
