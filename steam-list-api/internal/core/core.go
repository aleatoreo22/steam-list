package core

import (
	"steam-list-api.com/internal/database"
	"steam-list-api.com/internal/model"
	"steam-list-api.com/internal/service"
)

type Core struct {
	Token    Token
	Client   *service.Client
	Database *database.Database
}

type Token struct {
	IGDBClientID     string
	IGDBClientSecret string
	SteamwroksKey    string
}

func Initialize(igdbClientID string, igdbClientSecret string, steamworksKey string, dbType string) (*Core, error) {
	core := Core{}
	core.Token = Token{
		IGDBClientID:     igdbClientID,
		IGDBClientSecret: igdbClientSecret,
		SteamwroksKey:    steamworksKey,
	}
	service.CreateClient(core.Token.IGDBClientID, core.Token.IGDBClientSecret, core.Token.SteamwroksKey)
	db, err := database.Initialize(dbType)
	core.Database = db
	if err != nil {
		return nil, err
	}
	return &core, nil
}

func (core *Core) GetGame(id string) model.Game {
	game := core.Client.Game.GetGame(id)
	if game.IGDBID != 0 {
		//core.Database
	}
	return game
}
