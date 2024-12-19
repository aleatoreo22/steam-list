package database

import "steam-list-api.com/internal/model"

type GameDatabase struct {
	database *Database
}

func (gamedb *GameDatabase) Upsert(game model.Game) error {
	err := gamedb.database.db.Execute("")
	return err
}
