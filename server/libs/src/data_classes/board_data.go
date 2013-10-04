package data_classes

import (
	"time"
)

/* ---------------------- Board Data ---------------------- */

type BoardData struct {
	Id, UserId, GameId, Name string
	State, Active            int64
	DateAdded                time.Time
}
