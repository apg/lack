//go:generate -command yacc go tool yacc
//go:generate yacc -o query.go -p "query" query.y

package query

import "errors"

const eof = 0

// Parse returns a parsed query
func Parse(q []byte) (Query, error) {
	if queryParse(&queryLex{line: q}) == 0 {
		return lastQuery, nil
	}
	return nil, errors.New("Parse error!")
}
