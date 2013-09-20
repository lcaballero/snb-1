UpdateStream
	state (int)
	active (bool) // simulate deleted
	last-updated-by (fk)
	last-updated-on (time-stamp)

User
	id (pk)
	username
	password

UserInfo : UpdateStream
	id (pk)
	first-name
	last-name
	email
	display-name

UserToGroup
	id (pk)
	user-id
	group-id

SocialGroup
	id (pk)
	owner-id (fk)
	name

Board
	id (pk)
	game-id
	user-id
	name

Game
	id (pk)
	winning-board-id
	group-id
	sponsor-id
	active-state

GameTile
	id (pk)
	board-id
	criteria-id
	tile-position

Criteria : UpdateStream
	id (pk)
	description

Tags
	id (pk)
	asset-id (fk)
	description

Criteria Results
	id (pk)
	criteria-id
	user-id
	photo-id
	state // review, accepted, rejected



