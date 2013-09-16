package data_classes

import(
	"time"
)

/* ---------------------- Game Data ---------------------- */

type GameData struct {
	Anchor
}

func (game GameData) Id() string {
	return game.GetProp("id").(string)
}

func (game GameData) GroupId() string {
	return game.GetProp("group_id").(string)
}

func (game GameData) WinningBroardId() string {
	return game.GetProp("winning_board_id").(string)
}

func (game GameData) Name() string {
	return game.GetProp("name").(string)
}

func (game GameData) Description() string {
	return game.GetProp("description").(string)
}

func (game GameData) Active() bool {
	return game.GetProp("active").(bool)
}

func (game GameData) DateAdded() time.Time {
	return game.GetProp("date_added").(time.Time)
}