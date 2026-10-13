package service

import (
	"steam-list-api.com/internal/database"
	"steam-list-api.com/internal/thirdparty"
)

type Service struct {
	Game         *GameService
	thirdParty   *thirdparty.ThirdParty
	database     *database.Database
	itemsPerPage int
}

func CreateService(thirdParty *thirdparty.ThirdParty) *Service {
	client := &Service{
		thirdParty:   thirdParty,
		itemsPerPage: 20,
	}
	client.Game = &GameService{service: client}
	return client
}
