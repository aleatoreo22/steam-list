package sqlite

import (
	"database/sql"
	"errors"
	"os"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"steam-list-api.com/internal/model"
)

type SQLite struct {
	db *sql.DB
}

const sqliteDatabaseFile = "./steamlist.db"

func Initialize() (*SQLite, error) {
	sqlite := &SQLite{}
	err := sqlite.connect()
	if err != nil {
		return nil, err
	}
	defer sqlite.db.Close()
	return sqlite, nil
}

func (sqlite *SQLite) connect() error {
	db, err := sql.Open("sqlite3", sqliteDatabaseFile)
	if err != nil {
		return err
	}
	sqlite.db = db
	return nil
}

func (db *SQLite) DatabaseExists() bool {
	_, err := os.Stat(sqliteDatabaseFile)
	return !errors.Is(err, os.ErrNotExist)
}

func (sqlite *SQLite) createTable(model any) error {
	t := reflect.TypeOf(model)
	tableName := t.Name()
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("sql")
		columns = append(columns, tag)
	}
	sql := "CREATE TABLE IF NOT EXISTS " + tableName + " (" + strings.Join(columns, ",\n") + ");"
	_, err := sqlite.db.Exec(sql)
	return err
}

func (sqlite *SQLite) CreateDatabase() error {
	sqlite.connect()
	err := sqlite.createTable(model.Game{})
	if err != nil {
		return nil
	}
	defer sqlite.db.Close()
	return nil
}
