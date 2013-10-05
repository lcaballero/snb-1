CREATE TABLE IF NOT EXISTS UserToGroup
(
	id uuid NOT NULL,
	group_id uuid NOT NULL,
	user_id uuid NOT NULL,
	date_added timestamp without time zone,
	CONSTRAINT user_to_group_pkey PRIMARY KEY (id)
);