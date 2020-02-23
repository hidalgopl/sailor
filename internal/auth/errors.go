package auth

type authError struct {
	s string
}

func (e *authError) Error() string {
	return e.s
}

// New returns an error that formats as the given text.
func NewAuthError(text string) error {
	return &authError{text}
}
