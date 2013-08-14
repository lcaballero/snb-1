package session_manager

import (
	"fmt"
	"time"
)

type SessionError struct {
	When time.Time
	What string
}

func NewSessionError(format string, args ...interface{}) error {
	return SessionError{
		What: fmt.Sprintf(format, args...),
		When: time.Now(),
	}
}

func (e SessionError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}
