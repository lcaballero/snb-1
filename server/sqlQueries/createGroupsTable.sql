CREATE TABLE IF NOT EXISTS SocialGroup
(
	id uuid NOT NULL,
	group_name character varying(40) NOT NULL,
	group_desc text NOT NULL,
	group_owner uuid NOT NULL,
	date_added timestamp without time zone,
	CONSTRAINT social_group_pkey PRIMARY KEY (id)
);