insert into Game
	(id, group_id, name, description, state, active, date_added)
values
	($1, $2, $3, $4, $5, $6, now());