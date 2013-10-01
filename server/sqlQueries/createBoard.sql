INSERT INTO board
	(id, game_id, user_id, name, state, active, date_added)
VALUES
	($1, $2, $3, $4, $5, $6, now());