package core

import (
	"steam-list-api.com/internal/database"
	"steam-list-api.com/internal/service"
	"steam-list-api.com/internal/thirdparty"
	"steam-list-api.com/internal/thirdparty/tpgame"
	"steam-list-api.com/internal/thirdparty/tpstore"
	IGDBService "steam-list-api.com/pkg/igdb/igdbService"
	"steam-list-api.com/pkg/steamworks/steamworksService"
)

type Client struct {
	SteamworksClient *steamworksService.Client
	IGDBClient       *IGDBService.Client
}

type Core struct {
	Token    Token
	Service  *service.Service
	Database *database.Database
	Client   *Client
}

type Token struct {
	IGDBClientID     string
	IGDBClientSecret string
	SteamwroksKey    string
}

func CreateClient(clientIdIIGDB string, clientSecretIGDB string,
	keySteamwoks string) *Client {

	client := &Client{}
	client.SteamworksClient = steamworksService.NewClient(keySteamwoks)
	client.IGDBClient = IGDBService.NewClient(clientIdIIGDB, clientSecretIGDB)
	return client
}

func Initialize(igdbClientID string, igdbClientSecret string,
	steamworksKey string, dbType string) (*Core, error) {

	core := Core{}
	core.Client = CreateClient(igdbClientID, igdbClientSecret, steamworksKey)
	core.Token = Token{
		IGDBClientID:     igdbClientID,
		IGDBClientSecret: igdbClientSecret,
		SteamwroksKey:    steamworksKey,
	}

	var stores []tpstore.StoreThirdParty
	stores = append(stores, tpstore.NewTpSteam(core.Client.SteamworksClient))

	thirdParty := thirdparty.NewThirdParty(
		tpgame.NewTPGameIGDB(core.Client.IGDBClient),
		stores)

	core.Service = service.CreateService(thirdParty)
	db, err := database.Initialize(dbType)
	if err != nil {
		return nil, err
	}
	core.Database = db
	return &core, nil
}

// func (core *Core) GetPlayerGames(idPlayerSteam string, page int) ([]model.Game, error) {
// 	steamplayeridgames, err := core.service.game.getidsplayersteam(idplayersteam, page)
// 	if err != nil {
// 		return []model.game{}, err
// 	}
// 	var games []model.game
// 	for _, idgamesteam := range steamplayeridgames {
// 		game, err := core.database.game.getsteam(strconv.itoa(idgamesteam))
// 		if err != nil {
// 			return []model.game{}, err
// 		}
// 		if game.id != "" {
// 			games = append(games, game)
// 			continue
// 		}
// 		game = core.service.game.getsteam(idgamesteam)
// 		if game.igdbid == 0 {
// 			continue
// 		}
// 		game.steamappid = strconv.itoa(idgamesteam)
// 		games = append(games, game)
// 		err = core.database.game.upsert(game)
// 		if err != nil {
// 			log.error(err)
// 		}
// 	}
// 	return games, nil
// }
// func (core *Core) GetSteamIDByUsername(username string) {

// }
