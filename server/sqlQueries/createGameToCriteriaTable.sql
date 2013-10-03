CREATE TABLE gametocriteria
(
	id uuid NOT NULL,
	game_id uuid NOT NULL,
	criteria_id uuid NOT NULL,
	state INT,
	active INT,
	date_added timestamp without time zone,
	CONSTRAINT game_to_criteria_pkey PRIMARY KEY (id)
);