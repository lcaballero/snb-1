INSERT INTO tile
	(id, board_id, criteria_id, position, state, active, date_added)
VALUES
	($1, $2, $3, $4, $5, $6, now());