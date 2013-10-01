package data_classes

import (
	"time"
)

/* ---------------------- Board Data ---------------------- */

type BoardData struct {
	Anchor

	Id, UserId, GameId, Name string
	State, Active            int
	DateAdded                time.Time
}
