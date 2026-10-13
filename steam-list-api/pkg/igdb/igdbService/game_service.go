package igdbService

import (
	"encoding/json"
	"log"

	"steam-list-api.com/pkg/igdb/igdbModel"
)

type GameService struct {
	client *Client
}

func (service *GameService) GetExternal(query string) []igdbModel.Game {
	response := service.client.Post("external_games", query)
	var apiResponse []igdbModel.Game
	err := json.Unmarshal(response, &apiResponse)
	if err != nil {
		log.Fatalf("Erro ao deserializar o JSON: %v", err)
	}
	return apiResponse
}

func (service *GameService) Get(query string) []igdbModel.Game {
	response := service.client.Post("games", query)
	var apiResponse []igdbModel.Game
	err := json.Unmarshal(response, &apiResponse)
	if err != nil {
		log.Fatalf("Erro ao deserializar o JSON: %v", err)
	}
	return apiResponse
}
