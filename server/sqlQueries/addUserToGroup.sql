INSERT INTO usertogroup
	(id, group_id, user_id, date_added)
VALUES
	($1, $2, $3, now());