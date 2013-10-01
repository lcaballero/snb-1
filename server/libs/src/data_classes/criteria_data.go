package data_classes

import (
	"time"
)

/* ---------------------- Board Data ---------------------- */

type CriteriaData struct {
	Anchor

	Id, UserId, GameId, Name string
	State, Active            int
	DateAdded                time.Time
}
