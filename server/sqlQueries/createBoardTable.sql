CREATE TABLE board
(
	id uuid NOT NULL,
	game_id uuid NOT NULL,
	user_id uuid NOT NULL,
	name text,
	date_added timestamp without time zone,
	active int,
	CONSTRAINT board_pkey PRIMARY KEY (id)
);