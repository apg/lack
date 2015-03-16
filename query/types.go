package query

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/apg/lack/logfmt"
)

type operator int

const (
	Eq operator = iota
	Ne
	Ge
	Le
	Gt
	Lt
)

type Query interface {
	Match(*logfmt.Line) bool
}

type negQuery struct {
	q1 Query
}

type andQuery struct {
	q1 Query
	q2 Query
}

type orQuery struct {
	q1 Query
	q2 Query
}

type keyQuery struct {
	op  operator
	key string
	val interface{} // baseline to compare to.
}

type regexpQuery struct {
	re *regexp.Regexp
}

type inQuery struct {
	b []byte
}

func NewNegQuery(q Query) Query {
	return &negQuery{q}
}

func (q *negQuery) Match(line *logfmt.Line) bool {
	return !q.q1.Match(line)
}

func NewAndQuery(q1 Query, q2 Query) Query {
	return &andQuery{q1, q2}
}

func (q *andQuery) Match(line *logfmt.Line) bool {
	return q.q1.Match(line) && q.q2.Match(line)
}

func NewOrQuery(q1 Query, q2 Query) Query {
	return &orQuery{q1, q2}
}

func (q *orQuery) Match(line *logfmt.Line) bool {
	return q.q1.Match(line) || q.q2.Match(line)
}

func NewKeyQuery(op operator, key string, val interface{}) Query {
	return &keyQuery{op, key, val}
}

func (q *keyQuery) Match(line *logfmt.Line) bool {
	val, ok := line.Get(q.key)
	if !ok {
		return false
	}

	switch q.op {
	case Eq, Ne:
		eq := false

		// Lazily parse here.

		switch q.val.(type) {
		case *regexp.Regexp:
			re := q.val.(*regexp.Regexp)
			switch val.(type) {
			case int, int64:
				v := fmt.Sprintf("%d", val.(float64))
				eq = re.MatchString(v)
			case float64:
				v := fmt.Sprintf("%f", val.(float64))
				eq = re.MatchString(v)
			case string:
				eq = re.MatchString(val.(string))
			case []byte:
				eq = re.Match(val.([]byte))
			}
		case string:
			eq = val.(string) == q.val.(string)
		case []byte:
			eq = string(val.([]byte)) == string(q.val.([]byte))
		case int64, int:
			switch val.(type) {
			case int64, int:
				eq = val.(int64) == q.val.(int64)
			case []byte:
				if fix, err := strconv.ParseInt(string(val.([]byte)), 10, 64); err == nil {
					eq = fix == q.val.(int64)
				}
			}
		case float64:
			switch val.(type) {
			case float64:
				eq = val.(float64) == q.val.(float64)
			case []byte:
				if flo, err := strconv.ParseFloat(string(val.([]byte)), 64); err == nil {
					eq = flo == q.val.(float64)
				}
			}
		}

		if eq && q.op == Ne {
			return false
		} else if !eq && q.op == Ne {
			return true
		}
		return eq

	case Le, Lt, Ge, Gt:
		var l float64
		var r float64

		switch q.val.(type) {
		case float64:
			r = q.val.(float64)
		case int, int64:
			r = float64(q.val.(int64))
		default:
			return false
		}

		switch val.(type) {
		case float64:
			l = val.(float64)
		case int, int64:
			l = float64(val.(int64))
		case []byte:
			if flo, err := strconv.ParseFloat(string(val.([]byte)), 64); err != nil {
				return false
			} else {
				l = flo
			}
		case string:
			if flo, err := strconv.ParseFloat(val.(string), 64); err != nil {
				return false
			} else {
				l = flo
			}
		default:
			return false
		}

		switch q.op {
		case Le:
			return l <= r
		case Lt:
			return l < r
		case Ge:
			return l >= r
		case Gt:
			return l > r
		}
	}

	return false
}

func NewRegexpQuery(re *regexp.Regexp) Query {
	return &regexpQuery{re}
}

func (q *regexpQuery) Match(line *logfmt.Line) bool {
	return q.re.Match(line.Bytes())
}

func NewInQuery(s string) Query {
	return &inQuery{[]byte(s)}
}

func (q *inQuery) Match(line *logfmt.Line) bool {
	return bytes.Contains(line.Bytes(), q.b)
}
