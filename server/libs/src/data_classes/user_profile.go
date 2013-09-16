package data_classes

import(
	"time"
)

/* ---------------------- User Profile ---------------------- */

type UserProfile struct {
	Anchor
}

func (userProfile UserProfile) PrintAll() map[string]interface{} {
	m := make(map[string]interface{})

	m["Email"] = userProfile.Email()
	m["Date_Added"] = userProfile.DateAdded()

	return m
}

func (userProfile UserProfile) Id() string {
	return userProfile.GetProp("id").(string)
}

func (userProfile UserProfile) Email() string {
	return userProfile.GetProp("email").(string)
}

func (userProfile UserProfile) DateAdded() time.Time {
	return userProfile.GetProp("date_added").(time.Time)
}