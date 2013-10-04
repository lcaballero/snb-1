package data_classes

import (
	"time"
)

type UserProfile struct {
	Email, Id, Password string
	DateAdded           time.Time
}
