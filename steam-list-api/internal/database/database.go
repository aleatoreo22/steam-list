package database

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"steam-list-api.com/internal/database/sqlite"
	"steam-list-api.com/internal/model"
)

type Database struct {
	db   DatabaseFramework
	Game *GameDatabase
}

type DatabaseFramework interface {
	Open() error
	Close() error
	DatabaseExists() bool
	Execute(string) error
	CreateTable(any) error
	Query(string, any) ([]any, error)
}

type QueryBuilder struct {
	strings.Builder
}

func (qb *QueryBuilder) AddParameter(parameter string, data any) {
	typeName := reflect.TypeOf(data).Name()
	var dataString string
	switch typeName {
	case "string":
		dataString = "'" + data.(string) + "'"
	default:
		dataString = fmt.Sprintf("%v", data)
	}
	query := strings.Replace(qb.String(), parameter, dataString, 1)
	qb.Reset()
	qb.WriteString(query)
}

func Initialize(dbType string) (*Database, error) {
	var db DatabaseFramework
	switch dbType {
	case "sqlite":
		db = &sqlite.SQLite{}
	default:
		return nil, errors.New("unsupported database type")
	}
	if !db.DatabaseExists() {
		err := CreateDatabase(&db)
		if err != nil {
			return nil, err
		}
	}
	database := &Database{db: db}
	database.Game = &GameDatabase{database: database}
	return database, nil
}

func CreateDatabase(db *DatabaseFramework) error {
	err := (*db).Open()
	if err != nil {
		return err
	}
	err = (*db).CreateTable(model.Game{})
	if err != nil {
		return err
	}
	err = (*db).CreateTable(model.User{})
	if err != nil {
		return err
	}
	err = (*db).CreateTable(model.UserGame{})
	if err != nil {
		return err
	}
	(*db).Close()
	return nil
}
