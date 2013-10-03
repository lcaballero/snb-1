INSERT INTO gametocriteria
	(id, game_id, criteria_id, state, active, date_added)
VALUES
	($1, $2, $3, $4, $5, now());