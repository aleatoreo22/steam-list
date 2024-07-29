package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahmetb/go-linq/v3"
	"steam-list-api.com/internal/model"
	IGDB "steam-list-api.com/pkg/igdb"
	IGDBmodel "steam-list-api.com/pkg/igdb/igdbModel"
)

type GameService struct {
	client *Client
}

const itensPerPage = 10

func (service *GameService) GetTrendGames(page int) []model.Game {
	gamesIGDB := service.client.IGDBClient.Game.Get(IGDB.Fields("*") + IGDB.Sort("total_rating", true))
	var artworkIds []string
	linq.From(gamesIGDB).WhereT(
		func(game IGDBmodel.Game) bool {
			return len(game.Artworks) > 0
		}).SelectT(
		func(game IGDBmodel.Game) string {
			return strconv.Itoa(game.Artworks[0])
		}).ToSlice(&artworkIds)
	artworksIGDB := service.client.IGDBClient.ArtWork.Get(IGDB.Fields("*") + IGDB.Where(" id = ("+strings.Join(artworkIds, ",")+")"))
	coversIGDB := service.client.IGDBClient.Cover.Get(IGDB.Fields("*") + IGDB.Where(" id = ("+strings.Join(artworkIds, ",")+")"))
	games := convertGameIGDBToGame(gamesIGDB, artworksIGDB, coversIGDB)
	return games
}

func (service *GameService) getSteamGameIGDBId(id string) string {
	gameIGDB := service.client.IGDBClient.Game.GetExternal(IGDB.Fields("*") + IGDB.Where(" uid =\""+id+"\" & category = 1 ")) //Criar enumerado das plataformas
	//External Game Enums
	if len(gameIGDB) == 0 || gameIGDB[0].Game == 0 {
		return ""
	}
	idIGDB := strconv.Itoa(gameIGDB[0].Game)
	return idIGDB
}

func (service *GameService) GetPlayerGames(idPlayer string, page int) *[]model.Game {
	steamGames := service.client.SteamworksClient.Player.GetAllGames(idPlayer)
	var games []model.Game
	lastIndex := (itensPerPage * page)
	for i := lastIndex - itensPerPage; i < lastIndex; i++ {
		steamGame := steamGames.Games[i]
		idSteam := strconv.Itoa(steamGame.Appid)
		idIGDB := service.getSteamGameIGDBId(idSteam)
		if idIGDB == "" {
			continue
		}
		games = append(games, service.GetGame(idIGDB))
	}
	return &games
}

func (service *GameService) GetGame(id string) model.Game {
	gamesIGDB := service.client.IGDBClient.Game.Get(IGDB.Fields("*") + IGDB.Where(" id = "+id))
	if len(gamesIGDB) == 0 {
		return model.Game{}
	}
	artworksIGDB := []IGDBmodel.Artwork{}
	if len(gamesIGDB[0].Artworks) > 0 {
		artworkID := strconv.Itoa(gamesIGDB[0].Artworks[0])
		artworksIGDB = service.client.IGDBClient.ArtWork.Get(IGDB.Fields("*") + IGDB.Where(" id = "+artworkID))
	}
	coverIGDB := []IGDBmodel.Cover{}
	if gamesIGDB[0].Cover > 0 {
		coverId := strconv.Itoa(gamesIGDB[0].Cover)
		coverIGDB = service.client.IGDBClient.Cover.Get(IGDB.Fields("*") + IGDB.Where(" id = "+coverId))
	}
	games := convertGameIGDBToGame(gamesIGDB, artworksIGDB, coverIGDB)
	return games[0]
}

func convertGameIGDBToGame(gamesIGDB []IGDBmodel.Game, artworkIGDB []IGDBmodel.Artwork, coverIGDB []IGDBmodel.Cover) []model.Game {
	var games []model.Game
	linq.From(gamesIGDB).SelectT(
		func(g IGDBmodel.Game) model.Game {
			return model.Game{
				IGDBID: g.ID,
				Name:   g.Name,
			}
		}).ToSlice(&games)
	for i := 0; i < len(games); i++ {
		//Artwork
		artworkurl := linq.From(artworkIGDB).WhereT(func(artwork IGDBmodel.Artwork) bool {
			return artwork.Game == games[i].IGDBID
		}).SelectT(func(artwork IGDBmodel.Artwork) string {
			return artwork.URL
		}).First()
		if artworkurl != nil {
			games[i].ArtworkHDURL = fmt.Sprintf("%v", artworkurl)
			games[i].ArtworkHDURL = IGDB.SetImageSize(games[i].ArtworkHDURL, IGDB.HD)
		}
		//Cover
		coverurl := linq.From(coverIGDB).WhereT(func(cover IGDBmodel.Cover) bool {
			return cover.Game == games[i].IGDBID
		}).SelectT(func(cover IGDBmodel.Cover) string {
			return cover.URL
		}).First()
		if coverurl != nil {
			games[i].CoverHDURL = fmt.Sprintf("%v", coverurl)
			games[i].CoverHDURL = IGDB.SetImageSize(games[i].CoverHDURL, IGDB.CoverBig)
		}
	}
	return games
}
