package igdbService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	IGDBmodel "steam-list-api.com/pkg/igdb/igdbModel"
)

const baseurl = "https://api.igdb.com/v4/"

type Client struct {
	clientid     string
	clientsecret string
	token        string
	Game         GameService
	ArtWork      ArtworkService
	Cover        CoverService
}

func NewClient(clientid string, clientsecret string) *Client {
	client := &Client{
		clientid:     clientid,
		clientsecret: clientsecret,
	}
	client.Login()
	client.Game = GameService{client: client}
	client.ArtWork = ArtworkService{client: client}
	client.Cover = CoverService{client: client}
	return client
}

func (client *Client) Post(endpoint string, body string) []byte {
	headers := headerGeneration(client.clientid, client.token)
	requesturl := baseurl + endpoint
	req, err := http.NewRequest("POST", requesturl, bytes.NewReader([]byte(body)))
	if err != nil {
		log.Printf("%v", err)
	}
	for key, value := range headers {
		req.Header.Add(key, value)
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

func (client *Client) Login() {
	params := url.Values{}
	params.Add("client_id", client.clientid)
	params.Add("client_secret", client.clientsecret)
	params.Add("grant_type", "client_credentials")
	requesturl := fmt.Sprintf("%s?%s", "https://id.twitch.tv/oauth2/token", params.Encode())
	resp, err := http.Post(requesturl, "", nil)
	if err != nil {
		log.Printf("%v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler o corpo da resposta: %v", err)
	}
	var apiResponse IGDBmodel.Login
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Printf("Erro ao deserializar o JSON: %v", err)
	}
	if apiResponse.AccessToken == "" {
		log.Printf("Can't get IGDB Acess Token")
	}
	client.token = apiResponse.AccessToken
}

func headerGeneration(clientid string, token string) map[string]string {
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Authorization"] = "Bearer " + token
	headers["Client-ID"] = clientid
	return headers
}
