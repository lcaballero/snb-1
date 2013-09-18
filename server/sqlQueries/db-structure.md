User
	id (pk)
	username
	password

UserInfo
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
	owner-id
	name

Board
	id (pk)
	game-id
	user-id
	name
