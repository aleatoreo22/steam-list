package database

import (
	"steam-list-api.com/internal/database/sqlite"
)

type Database struct {
	sqlLite *sqlite.SQLite
}

func CreateConnection(connection string) (*Database, error) {
	db, err := sqlite.Initialize()
	if err != nil {
		return nil, err
	}
	if !db.DatabaseExists() {
		err = db.CreateDatabase()
		if err != nil {
			return nil, err
		}
	}
	return &Database{sqlLite: db}, nil
}

// func ConnectDatabase(connection string) (*sql.DB, error) {
// 	db, err := sql.Open("sqlite3", "./steamlist.db")
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	db.SetConnMaxLifetime(time.Minute * 3)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(10)
// 	rows, err := db.Query("SELECT IGDBID, ArtworkHDURL, CoverHDURL, Name, SteamAPPID FROM  game g ")
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}

// 	game := &model.Game{}
// 	columns, err := rows.Columns()

// 	obj := make([]interface{}, len(columns))
// 	for i := range obj {
// 		obj[i] = new(interface{})
// 	}

// 	for rows.Next() {
// 		err := rows.Scan(obj...)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	for i := 0; i < len(columns); i++ {
// 		structValue := reflect.ValueOf(game).Elem()
// 		fieldVal := structValue.FieldByName(columns[i])
// 		tete := obj[i]
// 		b, ok := tete.(*interface{})
// 		if !ok {

// 		}
// 		var v interface{}
// 		v = *b

// 		val := reflect.ValueOf(v)
// 		fieldVal.Set(val.Convert(fieldVal.Type()))
// 	}

// 	rows.Close()
// 	db.Close()
// 	return db, err
// }

// func convertValue(value interface{}) interface{} {
// 	v := reflect.ValueOf(value)

// 	switch v.Kind() {
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		return v.Int()
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 		return v.Uint()
// 	case reflect.Float32, reflect.Float64:
// 		return v.Float()
// 	case reflect.String:
// 		return v.String()
// 	case reflect.Bool:
// 		return v.Bool()
// 	default:
// 		return value
// 	}
// }
