package database

import (
	"steam-list-api.com/internal/model"
)

type GameDatabase struct {
	database *Database
}

func (gamedb *GameDatabase) get(query QueryBuilder) (model.Game, error) {
	err := gamedb.database.db.Open()
	if err != nil {
		return model.Game{}, err
	}
	result, err := gamedb.database.db.Query(query.String(), model.Game{})
	gamedb.database.db.Close()
	if err != nil {
		return model.Game{}, err
	}
	game := model.Game{}
	if len(result) > 0 {
		game = result[0].(model.Game)
	}
	return game, nil
}

func (gamedb *GameDatabase) GetIGDB(IGDBID string) (model.Game, error) {
	var query QueryBuilder
	query.WriteString("SELECT * FROM Game WHERE IGDBID = ?IGDBID")
	query.AddParameter("?IGDBID", IGDBID)
	return gamedb.get(query)
}

func (gamedb *GameDatabase) Get(id string) (model.Game, error) {
	var query QueryBuilder
	query.WriteString("SELECT * FROM Game WHERE id = ?id")
	query.AddParameter("?id", id)
	return gamedb.get(query)
}

func (gamedb *GameDatabase) Upsert(game model.Game) error {
	gameExists, err := gamedb.Get(game.ID)
	if err != nil {
		return err
	}
	var query QueryBuilder
	if gameExists.ID == "" {
		query.WriteString("INSERT INTO Game (id, igdbid, name, artworkhdurl, coverhdurl, steamappid) VALUES ")
		query.WriteString("( ?id, ?igdbid, ?name, ?artworkhdurl, ?coverhdurl, ?steamappid )")
	} else {
		query.WriteString("UPDATE Game SET igdbid = ?igdbid, name = ?name, artworkhdurl = ?artworkhdurl, ")
		query.WriteString("coverhdurl = ?coverhdurl, steamappid = ?steamappid WHERE id = ?id")
	}
	query.AddParameter("?id", game.ID)
	query.AddParameter("?igdbid", game.IGDBID)
	query.AddParameter("?name", game.Name)
	query.AddParameter("?artworkhdurl", game.ArtworkHDURL)
	query.AddParameter("?coverhdurl", game.CoverHDURL)
	query.AddParameter("?steamappid", game.SteamAPPID)
	err = gamedb.database.db.Open()
	if err != nil {
		return err
	}
	err = gamedb.database.db.Execute(query.String())
	gamedb.database.db.Close()
	return err
}
