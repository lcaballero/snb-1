CREATE TABLE IF NOT EXISTS criteria
(
	id uuid NOT NULL,
	description text,
	date_added timestamp without time zone,
	state int,
	active int,
	CONSTRAINT criteria_pkey PRIMARY KEY (id)
)