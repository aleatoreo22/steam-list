package model

type UserGame struct {
	UserID string `sql:"key srtlen 36"`
	GameID string `sql:"key srtlen 36"`
}
