package sqlite

import (
	"database/sql"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"steam-list-api.com/internal/model"
)

type SQLite struct {
	db *sql.DB
}

const sqliteDatabaseFile = "./steamlist.db"

func (sqlite *SQLite) Open() error {
	db, err := sql.Open("sqlite3", sqliteDatabaseFile)
	if err != nil {
		return err
	}
	sqlite.db = db
	return nil
}

func (sqlite *SQLite) Close() error {
	return sqlite.db.Close()
}

func (db *SQLite) DatabaseExists() bool {
	_, err := os.Stat(sqliteDatabaseFile)
	return !errors.Is(err, os.ErrNotExist)
}

func relativeType(typeName string, len int) string {
	relativeType := ""
	switch typeName {
	case "string":
		if len > 0 {
			relativeType = "varchar"
		} else {
			relativeType = "TEXT"
		}
	case "int", "int32", "int64":
		relativeType = "INTEGER"
	case "float32", "float64":
		relativeType = "REAL"
	case "bool":
		relativeType = "BOOLEAN"
	default:
		relativeType = "TEXT"
	}
	return relativeType
}

func (sqlite *SQLite) CreateTable(tableModel any) error {
	t := reflect.TypeOf(tableModel)
	tableName := t.Name()
	var columns []string
	var primaryKey []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("sql")
		var queryField strings.Builder
		queryField.WriteString(field.Name + " ")
		fieldLen := model.GetStringLen(tag)
		queryField.WriteString(relativeType(field.Type.Name(), fieldLen) + " ")
		if fieldLen > 0 {
			queryField.WriteString("(" + strconv.Itoa(fieldLen) + ")")
		}
		if model.IsPrimaryKey(tag) {
			queryField.WriteString(" NOT NULL ")
			primaryKey = append(primaryKey, field.Name)
		}
		columns = append(columns, queryField.String())
	}
	sql := "CREATE TABLE IF NOT EXISTS " + tableName + " (" + strings.Join(columns, ",\n")
	if len(primaryKey) > 0 {
		sql += ", \n PRIMARY KEY (" + strings.Join(primaryKey, ", ") + ")"
	}
	sql += ");"
	err := sqlite.Execute(sql)
	return err
}

func (sqlite *SQLite) Execute(query string) error {
	_, err := sqlite.db.Exec(query)
	return err
}

func (sqlite *SQLite) Query(query string, model any) ([]any, error) {
	rows, err := sqlite.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var results []any
	for rows.Next() {
		element := reflect.New(t).Elem()
		fields := make([]interface{}, len(columns))
		for i, column := range columns {
			field := element.FieldByName(column)
			if field.IsValid() {
				fields[i] = field.Addr().Interface()
			}
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, err
		}
		results = append(results, element.Interface())
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
