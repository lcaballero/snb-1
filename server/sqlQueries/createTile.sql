INSERT INTO tile
	(id, board_id, criteria_id, position, active, date_added)
VALUES
	($1, $2, $3, $4, $5, now());