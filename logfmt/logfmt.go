package logfmt

// Line represents a log line which may or may not have key-value pairs in it.
type Line struct {
	line []byte
	keys map[string]interface{}
}

// Bytes returns the raw log line
func (l *Line) Bytes() []byte {
	return l.line
}

// Get returns the key `key` if available in the Line.
func (l *Line) Get(key string) (interface{}, bool) {
	val, ok := l.keys[key]
	return val, ok
}
