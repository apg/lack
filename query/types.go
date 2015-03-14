package query

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"

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

type Matcher func(*logfmt.Line) bool

func NewNegMatcher(q Matcher) Matcher {
	return func(line *logfmt.Line) bool {
		return !q(line)
	}
}

func NewAndMatcher(q1 Matcher, q2 Matcher) Matcher {
	return func(line *logfmt.Line) bool {
		return q1(line) && q2(line)
	}
}

func NewOrMatcher(q1 Matcher, q2 Matcher) Matcher {
	return func(line *logfmt.Line) bool {
		return q1(line) && q2(line)
	}
}

func NewKeyMatcher(op operator, key string, mval interface{}) Matcher {
	return func(line *logfmt.Line) bool {
		val, ok := line.Get(key)
		if !ok {
			return false
		}

		switch op {
		case Eq, Ne:
			eq := false
			switch mval.(type) {
			case *regexp.Regexp:
				re := mval.(*regexp.Regexp)
				switch val.(type) {
				case int, int64:
					v := fmt.Sprintf("%d", val.(float64))
					eq = re.MatchString(v)
				case float64:
					v := fmt.Sprintf("%f", val.(float64))
					eq = re.MatchString(v)
				case string:
					eq = re.MatchString(val.(string))
				}
			default:
				eq = reflect.DeepEqual(val, mval)
			}

			if eq && op == Ne {
				return false
			} else if !eq && op == Ne {
				return true
			}

			return eq
		case Le, Lt, Ge, Gt:
			var l float64
			var r float64

			switch mval.(type) {
			case float64:
				r = mval.(float64)
			case int, int64:
				r = float64(mval.(int64))
			default:
				return false
			}

			switch val.(type) {
			case float64:
				l = val.(float64)
			case int, int64:
				l = float64(val.(int64))
			default:
				return false
			}

			switch op {
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
}

func NewRegexpMatcher(re *regexp.Regexp) Matcher {
	return func(line *logfmt.Line) bool {
		return re.Match(line.Bytes())
	}
}

func NewInMatcher(s string) Matcher {
	return func(line *logfmt.Line) bool {
		return bytes.Contains(line.Bytes(), []byte(s))
	}
}
