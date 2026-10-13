package model

type Game struct {
	ID           string `sql:"key strlen 36"`
	IGDBID       int    `json:"igdbid"`
	Name         string `sql:"strlen 255" json:"name"`
	ArtworkHDURL string `sql:"strlen 255" json:"artwork_hd_url"`
	CoverHDURL   string `sql:"strlen 255" json:"cover_hd_url"`
	SteamAPPID   string `sql:"strlen 255" json:"steam_appid"`
}
