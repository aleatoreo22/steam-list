package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"steam-list-api.com/internal/model"
	"steam-list-api.com/internal/service"
)

func main() {
	fmt.Println("Hello World")
	clientIdIGDB := loadEnv("IGDBClientID")
	clientSecretIGDB := loadEnv("IGDBClientSecret")
	keySteamworks := loadEnv("SteamwroksKey")
	router := mux.NewRouter()
	client := service.CreateClient(clientIdIGDB, clientSecretIGDB, keySteamworks)
	connectDatabase()

	router.HandleFunc("/api/hello",
		func(responseWriter http.ResponseWriter, request *http.Request) {
			response := model.APIResponse{Message: "Hello, World!"}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(response)
		}).Methods("GET")

	router.HandleFunc("/api/game/trend",
		func(responseWriter http.ResponseWriter, request *http.Request) {
			page, err := getPage(request)
			if err != nil {
				internalServerErrorHandler(responseWriter, request, err)
			}
			response := client.Game.GetTrendGames(page)
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(response)
		}).Methods("GET")

	router.HandleFunc("/api/game/{id}",
		func(responseWriter http.ResponseWriter, request *http.Request) {
			id := strings.TrimPrefix(request.URL.Path, "/api/game/")
			response := client.Game.GetGame(id)
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(response)
		}).Methods("GET")

	router.HandleFunc("/api/player/games/{id}",
		func(responseWriter http.ResponseWriter, request *http.Request) {
			id := strings.TrimPrefix(request.URL.Path, "/api/player/games/")
			page, err := getPage(request)
			if err != nil {
				internalServerErrorHandler(responseWriter, request, err)
			}
			response := client.Game.GetPlayerGames(id, page)
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(response)
		}).Methods("GET")
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getPage(r *http.Request) (int, error) {
	pageString := r.URL.Query().Get("page")
	if pageString == "" {
		return 1, nil
	}
	return strconv.Atoi(pageString)
}

func loadEnv(config string) string {
	file, err := os.Getwd()
	if err != nil {
		fmt.Println("Error to read file:", err)
	}
	file += "/.env"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error to read file:", err)
		return ""
	}
	tokens := strings.Split(string(content), "\n")
	token := ""
	for _, item := range tokens {
		if strings.Contains(item, config+"=") {
			token = strings.ReplaceAll(item, config+"=", "")
			break
		}
	}
	if token == "" {
		fmt.Println("Can't fount token " + config + "!")
	}
	return token
}

func internalServerErrorHandler(w http.ResponseWriter, _ *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error" + err.Error()))
}

func connectDatabase() (*sql.DB, error) {
	connectionString := loadEnv("MySQL")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	rows, err := db.Query("SELECT IGDBID, ArtworkHDURL, CoverHDURL, Name, SteamAPPID FROM  game g ")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	game := &model.Game{}
	columns, err := rows.Columns()

	obj := make([]interface{}, len(columns))
	for i := range obj {
		obj[i] = new(interface{})
	}

	for rows.Next() {
		err := rows.Scan(obj...)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < len(columns); i++ {
		structValue := reflect.ValueOf(game).Elem()
		fieldVal := structValue.FieldByName(columns[i])
		tete := obj[i]
		b, ok := tete.(*interface{})
		if !ok {

		}
		var v interface{}
		v = *b

		val := reflect.ValueOf(v)
		fieldVal.Set(val.Convert(fieldVal.Type()))
	}

	rows.Close()
	db.Close()
	return db, err
}

func convertValue(value interface{}) interface{} {
    v := reflect.ValueOf(value)

    switch v.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return v.Int()
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return v.Uint()
    case reflect.Float32, reflect.Float64:
        return v.Float()
    case reflect.String:
        return v.String()
    case reflect.Bool:
        return v.Bool()
    default:
        return value
    }
}
