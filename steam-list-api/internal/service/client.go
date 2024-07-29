package service

import (
	IGDBService "steam-list-api.com/pkg/igdb/igdbService"
	"steam-list-api.com/pkg/steamworks/steamworksService"
)

type Client struct {
	SteamworksClient *steamworksService.Client
	IGDBClient       *IGDBService.Client
	Game             *GameService
}

func CreateClient(clientIdIIGDB string, clientSecretIGDB string, keySteamwoks string) *Client {
	client := &Client{}
	client.Game = &GameService{client: client}
	client.IGDBClient = IGDBService.NewClient(clientIdIIGDB, clientSecretIGDB)
	client.SteamworksClient = steamworksService.CreateClient(keySteamwoks)
	return client
}
