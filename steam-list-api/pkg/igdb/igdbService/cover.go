package igdbService

import (
	"encoding/json"
	"log"

	IGDBmodel "steam-list-api.com/pkg/igdb/igdbModel"
)

type CoverService struct {
	client *Client
}

func (service *CoverService) Get(query string) []IGDBmodel.Cover{
	response := service.client.Post("covers", query)
	var apiResponse []IGDBmodel.Cover
	err := json.Unmarshal(response, &apiResponse)
	if err != nil {
		log.Fatalf("Erro ao deserializar o JSON: %v", err)
	}
	return apiResponse
}
