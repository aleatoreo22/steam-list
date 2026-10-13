package tpstore

import (
	"strconv"

	"github.com/ahmetb/go-linq/v3"
	"steam-list-api.com/pkg/steamworks/steamworksModel"
	"steam-list-api.com/pkg/steamworks/steamworksService"
)

type TpSteam struct {
	client *steamworksService.Client
}

func NewTpSteam(client *steamworksService.Client) *TpSteam {
	return &TpSteam{client: client}
}

func (t *TpSteam) GetPlayerGames(playerID string, page int, itemsPerPage int) (*[]string, error) {
	games := t.client.Player.GetAllGames(playerID)

	lastIndex := (itemsPerPage * page)
	var gamesIds []int
	for i := lastIndex - itemsPerPage; i < lastIndex; i++ {
		if i > len(games.Games)-1 {
			break
		}
		steamGame := games.Games[i]
		gamesIds = append(gamesIds, steamGame.Appid)
	}

	var gameids []string

	linq.From(games.Games).SelectT(
		func(g steamworksModel.Game) string {
			return strconv.Itoa(g.Appid)
		}).ToSlice(gameids)

	return &gameids, nil
}

func (t *TpSteam) StoreID() string {
	return "steam"
}
