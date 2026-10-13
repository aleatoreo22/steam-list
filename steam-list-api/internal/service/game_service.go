package service

import (
	"steam-list-api.com/internal/model"
)

type GameService struct {
	service *Service
}

func (g *GameService) GetTrend(page int) *[]model.Game {
	games, err := g.service.thirdParty.GetTrend()
	if err != nil {
		return nil
	}
	return games
}

func (g *GameService) GetSteamPlayerGames(playerId string, page int) (*[]model.Game, error) {
	gamesID, err := g.service.thirdParty.GetPlayerGamesId(
		"steam", playerId, page, g.service.itemsPerPage)

	if err != nil {
		return nil, err
	}

	var games []model.Game
	for _, idgamesteam := range *gamesID {
		game, err := g.service.database.Game.GetSteam(idgamesteam)
		if err != nil {
			return nil, err
		}
		if game.ID != "" {
			games = append(games, *game)
			continue
		}
		if game.IGDBID == 0 {
			continue
		}
		game.SteamAPPID = idgamesteam
		games = append(games, *game)
		_ = g.service.database.Game.Upsert(*game)
	}
	return &games, nil

	// games, err := g.service.thirdParty.GetPlayerGames("steam", idPlayer,
	// 	page, g.service.itemsPerPage)
	// if err != nil {
	// 	return nil, nil
	// }
	// return games, nil
}

func (g *GameService) GetGameByID(id string) *model.Game {
	game, err := g.service.thirdParty.GetGameByID(id)
	if err != nil {
		return nil
	}
	return game
}
