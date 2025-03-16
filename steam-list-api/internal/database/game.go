package database

import (
	"github.com/google/uuid"
	"steam-list-api.com/internal/model"
)

type GameDatabase struct {
	database *Database
}

func (db *GameDatabase) get(query QueryBuilder) (model.Game, error) {
	err := db.database.db.Open()
	if err != nil {
		return model.Game{}, err
	}
	result, err := db.database.db.Query(query.String(), model.Game{})
	db.database.db.Close()
	if err != nil {
		return model.Game{}, err
	}
	game := model.Game{}
	if len(result) > 0 {
		game = result[0].(model.Game)
	}
	return game, nil
}

func (db *GameDatabase) GetIGDB(idIGDB string) (model.Game, error) {
	var query QueryBuilder
	query.WriteString("SELECT * FROM Game WHERE IGDBID = ?IGDBID")
	query.AddParameter("?IGDBID", idIGDB)
	return db.get(query)
}

func (db *GameDatabase) Get(id string) (model.Game, error) {
	var query QueryBuilder
	query.WriteString("SELECT * FROM Game WHERE id = ?id")
	query.AddParameter("?id", id)
	return db.get(query)
}

func (db *GameDatabase) verifyIfExists(game *model.Game) (bool, error) {
	gameExists, err := db.Get(game.ID)
	if err != nil {
		return false, err
	}
	if gameExists.ID == "" {
		// if game.IGDBID > 0 {
		// 	gameExists, err = db.GetIGDB(strconv.Itoa(game.IGDBID))
		// }
		// if err != nil {
		// 	return false, err
		// }
		// if gameExists.ID == "" {

		//Existem jogos iguais na steam com versao pra Mac e pra Windows/Linux,
		// eles sao 2 IDs na steam e apenas 1 no IGDB
		if game.SteamAPPID != "" {
			gameExists, err = db.GetSteam(game.SteamAPPID)
		}
		if err != nil {
			return false, err
		}
		// }
	}
	game = &gameExists
	return game.ID != "", nil
}

func (db *GameDatabase) Upsert(game model.Game) error {
	exist, err := db.verifyIfExists(&game)
	if err != nil {
		return err
	}
	var query QueryBuilder
	if !exist {
		game.ID = uuid.NewString()
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
	err = db.database.db.Open()
	if err != nil {
		return err
	}
	err = db.database.db.Execute(query.String())
	db.database.db.Close()
	return err
}

func (db *GameDatabase) GetSteam(idSteam string) (model.Game, error) {
	var query QueryBuilder
	query.WriteString("SELECT * FROM Game WHERE SteamAPPID = ?SteamAPPID")
	query.AddParameter("?SteamAPPID", idSteam)
	return db.get(query)
}
