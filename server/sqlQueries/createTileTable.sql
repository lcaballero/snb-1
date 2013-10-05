CREATE TABLE IF NOT EXISTS tile
(
	id uuid NOT NULL,
	board_id uuid NOT NULL,
	criteria_id uuid NOT NULL,
	position INTEGER,
	state INTEGER,
	active INTEGER,
	date_added timestamp without time zone,
	CONSTRAINT tile_pkey PRIMARY KEY (id)
)