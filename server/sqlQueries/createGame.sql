insert into Game
	(id, group_id, winning_board_id, name, description, sponsor_id, active, date_added)
values
	($1, $2, $3, $4, $5, $6, $7, now());