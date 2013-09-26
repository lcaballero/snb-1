INSERT INTO criteria
	(id, description, active, date_added)
VALUES
	($1, $2, $3, now());