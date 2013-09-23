INSERT INTO board
	(id, game_id, user_id, name, active, date_added)
VALUES
	($1, $2, $3, $4, $5, now());