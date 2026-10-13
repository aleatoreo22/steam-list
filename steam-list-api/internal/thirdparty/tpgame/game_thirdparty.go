package tpgame

import "steam-list-api.com/internal/model"

type GameThirdParty interface {
	GetTrend() (*[]model.Game, error)
	GetGameByStoreID(storeID string, storeType string) (*model.Game, error)
	GetGameByID(gameID string) (*model.Game, error)
}
