package igdbService

import (
	"encoding/json"
	"log"

	IGDBmodel "steam-list-api.com/pkg/igdb/igdbModel"
)

type GameService struct {
	client *Client
}

func (service *GameService) GetExternal(query string) []IGDBmodel.Game {
	response := service.client.Post("external_games", query)
	var apiResponse []IGDBmodel.Game
	err := json.Unmarshal(response, &apiResponse)
	if err != nil {
		log.Fatalf("Erro ao deserializar o JSON: %v", err)
	}
	return apiResponse
}

func (service *GameService) Get(query string) []IGDBmodel.Game {
	response := service.client.Post("games", query)
	var apiResponse []IGDBmodel.Game
	err := json.Unmarshal(response, &apiResponse)
	if err != nil {
		log.Fatalf("Erro ao deserializar o JSON: %v", err)
	}
	return apiResponse
}
