package data_classes

import (
	"time"
)

/* ---------------------- Board Data ---------------------- */

type TileData struct {
	Anchor

	Id, BoardId, CriteriaId string
	Position, Active, State int64
	DateAdded               time.Time
}
