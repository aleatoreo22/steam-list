package model

type Game struct {
	IGDBID       int    `sql:"id primary key" json:"igdbid"`
	Name         string `sql:"name text" json:"name"`
	ArtworkHDURL string `sql:"artwork_hd_url text" json:"artwork_hd_url"`
	CoverHDURL   string `sql:"cover_hd_url text" json:"cover_hd_url"`
	SteamAPPID   string `sql:"SteamAPPID text" json:"steam-appid"`
}
