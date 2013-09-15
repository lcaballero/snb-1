insert into _user
	(id, email, password, date_added)
values
	($1, $2, $3, now());