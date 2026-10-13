package thirdparty

import (
	"steam-list-api.com/internal/model"
	"steam-list-api.com/internal/thirdparty/tpgame"
	"steam-list-api.com/internal/thirdparty/tpstore"
)

type ThirdParty struct {
	gameThirdParty  tpgame.GameThirdParty
	storeThirdParty []tpstore.StoreThirdParty
}

func NewThirdParty(gameThirdParty tpgame.GameThirdParty,
	storeThirdParty []tpstore.StoreThirdParty) *ThirdParty {
	return &ThirdParty{
		gameThirdParty:  gameThirdParty,
		storeThirdParty: storeThirdParty,
	}
}

func (t *ThirdParty) GetTrend() (*[]model.Game, error) {
	return t.gameThirdParty.GetTrend()
}

func (t *ThirdParty) GetGameByStoreID(storeID string,
	storeType string) (*model.Game, error) {
	return t.gameThirdParty.GetGameByStoreID(storeID, storeType)
}

func (t *ThirdParty) GetGamesByStoreID(storeID string,
	gamesID []string) (*[]model.Game, error) {
	games := make([]model.Game, 0, len(gamesID))
	for _, id := range gamesID {
		game, err := t.gameThirdParty.GetGameByStoreID(id, storeID)
		if err != nil {
			return nil, err
		}
		games = append(games, *game)
	}
	return &games, nil
}

func (t *ThirdParty) GetPlayerGamesId(storeID string, playerID string,
	page int, itemsPerPage int) (*[]string, error) {

	for _, store := range t.storeThirdParty {
		if store.StoreID() == storeID {
			gamesID, err := store.GetPlayerGames(playerID, page, itemsPerPage)
			if err != nil {
				return nil, err
			}
			return gamesID, err
		}
	}

	//Error
	return nil, nil
}

func (t *ThirdParty) GetPlayerGames(storeID string, playerID string,
	page int, itemsPerPage int) (*[]model.Game, error) {

	for _, store := range t.storeThirdParty {
		if store.StoreID() == storeID {
			gamesID, err := store.GetPlayerGames(playerID, page, itemsPerPage)
			if err != nil {
				return nil, err
			}
			return t.GetGamesByStoreID(storeID, *gamesID)
		}
	}

	//Error
	return nil, nil
}

func (t *ThirdParty) GetGameByID(gameID string) (*model.Game, error) {
	return t.gameThirdParty.GetGameByID(gameID)
}
