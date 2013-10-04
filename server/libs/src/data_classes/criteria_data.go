package data_classes

import (
	"time"
)

/* ---------------------- Board Data ---------------------- */

type CriteriaData struct {
	Id, UserId, GameId, Name string
	State, Active            int
	DateAdded                time.Time
}
