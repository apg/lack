package logfmt

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
