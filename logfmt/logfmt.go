package logfmt

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// Line represents a log line which may or may not have key-value pairs in it.
type Line struct {
	line []byte
	keys []string
	vals []interface{}
}

// NewLine returns a new buffer
func NewLine() *Line {
	return &Line{
		line: make([]byte, 0),
		keys: make([]string, 0),
		vals: make([]interface{}, 0),
	}
}

// Bytes returns the raw log line
func (l *Line) Bytes() []byte {
	return l.line
}

// Append a field onto the lin
func (l *Line) Append(key string, val interface{}) {
	l.keys = append(l.keys, key)
	l.vals = append(l.vals, val)
}

// Get returns the key `key` if available in the Line.
func (l *Line) Get(key string) (interface{}, bool) {
	for i, vkey := range l.keys {
		if key == vkey {
			return l.vals[i], true
		}
	}
	return nil, false
}

// Reset sets line to line
func (l *Line) Reset(line []byte) {
	l.line = line
	l.keys = l.keys[0:0]
	l.vals = l.vals[0:0]
}

// Format a line for output
// %{key}, %1 ' ' separated field
func (l *Line) Format(format []byte) (string, error) {
	var splits [][]byte
	var output bytes.Buffer
	i := bytes.IndexByte(format, '%')
	for i != -1 {
		if i > 0 {
			//			log.Printf("Writing: %q, i=%d", string(format[0:i]), i)
			output.Write(format[0:i])
		}
		format = format[i+1:]

		if len(format) > 0 {
			switch format[0] {
			case '%':
				//				log.Printf("Writing: %%, i=%d", i)
				output.WriteByte('%')
				format = format[i+1:]
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// now we need to split the line on ' '
				var j int
				for j = 1; j < len(format) && format[j] >= '0' && format[j] <= '9'; j++ {
				}
				if splits == nil {
					splits = bytes.Split(l.line, []byte{' '})
				}
				if index, err := strconv.Atoi(string(format[0:j])); err != nil {
					//					log.Println("Number: ", string(format), "for: j=", j, string(format[0:j]))
					return "", err
				} else if index < len(splits) {
					if index > 0 {
						output.Write(splits[index-1])
					} else {
						output.Write(l.line)
					}
				}
				format = format[j:]
			case '{':
				keyIndex := bytes.IndexByte(format, '}')
				if keyIndex <= 0 {
					return "", errors.New("Bad key format")
				}
				key := format[1:keyIndex]
				//				log.Printf("Key: %q, keyIndex=%d, key=%q", string(format), keyIndex, string(format[1:keyIndex]))
				if v, ok := l.Get(string(key)); ok {
					output.WriteString(string(fmt.Sprintf("%s", v)))
				}
				format = format[keyIndex+1:]
			default:
				return "", errors.New("Bad key format")
			}
		}
		i = bytes.IndexByte(format, '%')
	}

	return output.String(), nil
}
