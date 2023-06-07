package errors

import (
	"errors"
	"fmt"
	"io"
)

// Formattable represents a formattable error.
type Formattable interface {
	error
	Message(verbose bool) string
}

// Format formats an error.
func Format(err Formattable, s fmt.State, verb rune) {
	switch {
	case verb == 'v' && s.Flag('+'):
		werr := errors.Unwrap(err)
		if werr != nil {
			// It's just a safety check.
			// Unwrap() should not return nil, but nothing forbids it.
			_, _ = fmt.Fprintf(s, "%+v", werr)
		}
		msg := err.Message(true)
		if msg != "" {
			_, _ = fmt.Fprintf(s, "\n%s", msg)
		}
	case verb == 'v' || verb == 's':
		_, _ = io.WriteString(s, err.Error())
	case verb == 'q':
		_, _ = fmt.Fprintf(s, "%q", err.Error())
	}
}

// Error formats an error on a single line.
func Error(err Formattable) string {
	msg := err.Message(false)
	var wmsg string
	werr := errors.Unwrap(err)
	if werr != nil {
		// It's just a safety check.
		// Unwrap() should not return nil, but nothing forbids it.
		wmsg = werr.Error()
	}
	if msg == "" {
		return wmsg
	}
	return msg + ": " + wmsg
}
