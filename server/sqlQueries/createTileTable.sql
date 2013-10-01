CREATE TABLE tile
(
	id uuid NOT NULL,
	board_id uuid NOT NULL,
	criteria_id uuid NOT NULL,
	position INT,
	active INT,
	date_added timestamp without time zone,
	CONSTRAINT tile_pkey PRIMARY KEY (id)
)