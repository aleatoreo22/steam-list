package igdbService

import (
	"encoding/json"
	"log"

	"steam-list-api.com/pkg/igdb/igdbModel"
)

type ArtworkService struct {
	client *Client
}

func (service *ArtworkService) Get(query string) []igdbModel.Artwork {
	response := service.client.Post("artworks", query)
	var apiResponse []igdbModel.Artwork
	err := json.Unmarshal(response, &apiResponse)
	if err != nil {
		log.Fatalf("Erro ao deserializar o JSON: %v", err)
	}
	return apiResponse
}
