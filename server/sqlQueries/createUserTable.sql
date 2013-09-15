CREATE TABLE _User (
	id uuid NOT NULL,
	Email varchar(255),
	Password text NOT NULL,
	date_added timestamp without time zone,
	CONSTRAINT _user_pkey PRIMARY KEY (id)
);