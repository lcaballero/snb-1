package data_classes

import (
	"time"
)

/* ---------------------- Criteria Data ---------------------- */

type CriteriaData struct {
	Id, Description string
	State, Active   int64
	DateAdded       time.Time
}
