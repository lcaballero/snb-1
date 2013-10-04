package data_classes

import (
	"time"
)

/* ---------------------- Board Data ---------------------- */

type TileData struct {
	Id, BoardId, CriteriaId string
	Position, Active, State int64
	DateAdded               time.Time
}
