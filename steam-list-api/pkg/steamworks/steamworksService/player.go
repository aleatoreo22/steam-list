package steamworksService

import (
	"encoding/json"
	"log"

	steamworksModel "steam-list-api.com/pkg/steamworks/steamworksModel"
)

type PlayerService struct {
	client *Client
}

func (service *PlayerService) GetAllGames(accountId string) *steamworksModel.Games {
	response := service.client.Get("IPlayerService/GetOwnedGames/v1", accountId)
	var apiResponse steamworksModel.Root[steamworksModel.Games]
	err := json.Unmarshal(response, &apiResponse)
	if err != nil {
		log.Fatalf("Erro ao deserializar o JSON: %v", err)
	}
	return &apiResponse.Response
}
