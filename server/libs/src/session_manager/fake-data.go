package session_manager

func User1() *User {
	return &User{
		FirstName:  "Lucas",
		LastName:   "Caballero",
		UserName:   "caballero.luke@gmail.com",
		InternalId: "user-id-1",
		SessionId:  "",
		UserGroup:  "group-1",
		Teir:       "default-teir",
		Email:      "caballero.luke@gmail.com",
		Phone:      "555-555-5551",
		Address:    "555 West 5th Street, Fifth City, Fifth State 50505",
		Company:    "5th Dimension Company",
		Location:   "Home",
		Url:        "digital-lotion.com",
		Created_At: "now",
		Updated_At: "later",
		Updated_By: "user-1",
	}
}

func User2() *User {
	return &User{
		FirstName:  "Lucas",
		LastName:   "Caballero",
		UserName:   "caballero.luke@gmail.com",
		InternalId: "user-id-2",
		SessionId:  "",
		UserGroup:  "group-1",
		Teir:       "default-teir",
		Email:      "caballero.luke@gmail.com",
		Phone:      "555-555-5551",
		Address:    "555 West 5th Street, Fifth City, Fifth State 50505",
		Company:    "5th Dimension Company",
		Location:   "Home",
		Url:        "digital-lotion.com",
		Created_At: "now",
		Updated_At: "later",
		Updated_By: "user-1",
	}
}
