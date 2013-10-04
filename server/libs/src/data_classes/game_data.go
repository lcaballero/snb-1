package data_classes

import (
	"time"
)

/* ---------------------- Game Data ---------------------- */

type GameData struct {
	Id, GroupId, WinningBoardId, Name, Description string
	State                                          int64
	DateAdded                                      time.Time
}
