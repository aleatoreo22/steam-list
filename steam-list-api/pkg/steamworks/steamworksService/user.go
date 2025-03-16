package steamworksService

type UserService struct {
	client *Client
}

func (service *PlayerService) GetSteamIDByUsername(username string) {
	response := service.client.Get("ISteamUser/ResolveVanityURL/v1/", username)
	if response != nil {

	}
}

// https://api.steampowered.com/
