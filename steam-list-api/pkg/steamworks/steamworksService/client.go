package steamworksService

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const baseurl = "https://api.steampowered.com/"

type Client struct {
	key    string
	Player *PlayerService
}

func CreateClient(key string) *Client {
	client := &Client{
		key: key,
	}
	client.Player = &PlayerService{
		client: client,
	}
	return client
}

func (client *Client) Get(endpoint string, accountId string) []byte {
	var params = url.Values{}
	params.Add("key", client.key)
	params.Add("steamid", accountId)
	requesturl := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	requesturl = baseurl + requesturl
	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		log.Printf("%v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("%v", err)
	}
	defer resp.Body.Close()
	//TODO test error rest and show api message
	responsecontent, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler o corpo da resposta: %v", err)
	}
	return responsecontent
}
