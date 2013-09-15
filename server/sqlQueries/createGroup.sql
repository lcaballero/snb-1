insert into SocialGroup
	(id, group_name, group_desc, group_owner, date_added)
values
	($1, $2, $3, $4, now());