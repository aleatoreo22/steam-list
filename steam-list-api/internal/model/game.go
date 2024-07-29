package model

type Game struct {
	IGDBID       int    `json:"igdbid"`
	Name         string `json:"name"`
	ArtworkHDURL string `json:"artwork_hd_url"`
	CoverHDURL   string `json:"cover_hd_url"`
	SteamAPPID   string `json:"steam-appid"`
}
