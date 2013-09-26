CREATE TABLE criteria
(
	id uuid NOT NULL,
	description text,
	date_added timestamp without time zone,
	active int,
	CONSTRAINT criteria_pkey PRIMARY KEY (id)
)