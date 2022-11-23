package middleware

type ContextKey string

// String returns the string representation of the context key.
func (c ContextKey) String() string {
	return string(c)
}
