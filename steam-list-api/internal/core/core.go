package core

import (
	"strconv"

	"github.com/labstack/gommon/log"
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
	core.Client = service.CreateClient(core.Token.IGDBClientID, core.Token.IGDBClientSecret, core.Token.SteamwroksKey)
	db, err := database.Initialize(dbType)
	core.Database = db
	if err != nil {
		return nil, err
	}
	return &core, nil
}

func (core *Core) GetGame(id string) model.Game {
	game, err := core.Database.Game.GetIGDB(id)
	if err != nil {
		return model.Game{}
	}
	if game.ID != "" {
		return game
	}
	game = core.Client.Game.GetIGDB(id)
	if game.IGDBID != 0 {
		err := core.Database.Game.Upsert(game)
		if err != nil {
			log.Error(err)
		}
	}
	return game
}

func (core *Core) GetPlayerGames(idPlayerSteam string, page int) ([]model.Game, error) {
	steamPlayerIdGames, err := core.Client.Game.GetIdsPlayerSteam(idPlayerSteam, page)
	if err != nil {
		return []model.Game{}, err
	}
	var games []model.Game
	for _, idGameSteam := range steamPlayerIdGames {
		game, err := core.Database.Game.GetSteam(strconv.Itoa(idGameSteam))
		if err != nil {
			return []model.Game{}, err
		}
		if game.ID != "" {
			games = append(games, game)
			continue
		}
		game = core.Client.Game.GetSteam(idGameSteam)
		if game.IGDBID == 0 {
			continue
		}
		game.SteamAPPID = strconv.Itoa(idGameSteam)
		games = append(games, game)
		err = core.Database.Game.Upsert(game)
		if err != nil {
			log.Error(err)
		}
	}
	return games, nil
}
func (core *Core) GetSteamIDByUsername(username string) {

}
