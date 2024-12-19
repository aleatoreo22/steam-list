package model

type User struct {
	ID string `sql:"key strlen 36" json:"id"`
	SteamID string `sql:"strlen 17"`
}
