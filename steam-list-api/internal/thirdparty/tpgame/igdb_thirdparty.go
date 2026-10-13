package tpgame

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahmetb/go-linq/v3"
	"steam-list-api.com/internal/model"
	IGDB "steam-list-api.com/pkg/igdb"
	IGDBmodel "steam-list-api.com/pkg/igdb/igdbModel"
	IGDBService "steam-list-api.com/pkg/igdb/igdbService"
)

type TPGameIGDB struct {
	client *IGDBService.Client
}

func NewTPGameIGDB(client *IGDBService.Client) *TPGameIGDB {
	return &TPGameIGDB{client: client}
}

func (t *TPGameIGDB) GetTrend() (*[]model.Game, error) {
	gamesIGDB := t.client.Game.Get(IGDB.Fields("*") + IGDB.Sort("total_rating", true))
	var artworkIds []string
	linq.From(gamesIGDB).WhereT(
		func(game IGDBmodel.Game) bool {
			return len(game.Artworks) > 0
		}).SelectT(
		func(game IGDBmodel.Game) string {
			return strconv.Itoa(game.Artworks[0])
		}).ToSlice(&artworkIds)
	artworksIGDB := t.client.ArtWork.Get(IGDB.Fields("*") + IGDB.Where(" id = ("+strings.Join(artworkIds, ",")+")"))
	coversIGDB := t.client.Cover.Get(IGDB.Fields("*") + IGDB.Where(" id = ("+strings.Join(artworkIds, ",")+")"))
	return convertGameIGDBToGame(gamesIGDB, artworksIGDB, coversIGDB), nil
}

func (t *TPGameIGDB) GetGameByStoreID(storeID string, storeType string) (*model.Game, error) {
	var game *model.Game
	var err error
	err = nil

	switch storeType {
	case "steam":
		gameIGDB := t.client.Game.
			GetExternal(IGDB.Fields("*") + IGDB.Where(" uid =\""+storeID+"\" & category = 1 "))
		if len(gameIGDB) > 0 && gameIGDB[0].Game > 0 {
			game, err = t.GetGameByID(strconv.Itoa(gameIGDB[0].Game))
		}
	}

	return game, err
}

func (t *TPGameIGDB) GetGameByID(idIGDB string) (*model.Game, error) {
	gamesIGDB := t.client.Game.Get(IGDB.Fields("*") + IGDB.Where(" id = "+idIGDB))
	if len(gamesIGDB) == 0 {
		return &model.Game{}, nil
	}

	artworksIGDB := []IGDBmodel.Artwork{}
	if len(gamesIGDB[0].Artworks) > 0 {
		artworkID := strconv.Itoa(gamesIGDB[0].Artworks[0])
		artworksIGDB = t.client.ArtWork.Get(IGDB.Fields("*") + IGDB.Where(" id = "+artworkID))
	}

	coverIGDB := []IGDBmodel.Cover{}
	if gamesIGDB[0].Cover > 0 {
		coverId := strconv.Itoa(gamesIGDB[0].Cover)
		coverIGDB = t.client.Cover.Get(IGDB.Fields("*") + IGDB.Where(" id = "+coverId))
	}

	games := convertGameIGDBToGame(gamesIGDB, artworksIGDB, coverIGDB)
	return &(*games)[0], nil
}

func convertGameIGDBToGame(gamesIGDB []IGDBmodel.Game,
	artworkIGDB []IGDBmodel.Artwork,
	coverIGDB []IGDBmodel.Cover) *[]model.Game {

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
	return &games
}
