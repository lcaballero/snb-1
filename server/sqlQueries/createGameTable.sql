CREATE TABLE Game
(
	id uuid NOT NULL,
	group_id uuid NOT NULL,
	winning_board_id uuid,
	name character varying(40) NOT NULL,
	description text NOT NULL,
	sponsor_id uuid,
	state int,
	date_added timestamp without time zone,
	CONSTRAINT game_pkey PRIMARY KEY (id)
);