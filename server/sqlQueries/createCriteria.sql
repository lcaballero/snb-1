INSERT INTO criteria
	(id, description, state, active, date_added)
VALUES
	($1, $2, $3, $4, now());