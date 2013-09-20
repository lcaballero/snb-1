CREATE TABLE _User (
	id uuid NOT NULL,
	email varchar(255),
	password text NOT NULL,
	date_added timestamp without time zone,
	CONSTRAINT _user_pkey PRIMARY KEY (id)
);