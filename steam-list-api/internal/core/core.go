package core

import (
	"steam-list-api.com/internal/database"
	"steam-list-api.com/internal/service"
)

type Core struct {
	Token  Token
	Client *service.Client
}

type Token struct {
	IGDBClientID     string
	IGDBClientSecret string
	SteamwroksKey    string
	SQLite           string
}

func Initialize(igdbClientID string, igdbClientSecret string, steamworksKey string) *Core {
	core := Core{}
	// service.CreateClient(core.Token.IGDBClientID, core.Token.IGDBClientSecret, core.Token.SteamwroksKey)
	database.CreateConnection("")
	return &core
}
