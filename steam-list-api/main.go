package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"steam-list-api.com/internal/core"
	"steam-list-api.com/internal/model"
)

func main() {
	fmt.Println("Hello World")
	IGDBClientId := loadEnv("IGDBClientID")
	IGDBClientSecret := loadEnv("IGDBClientSecret")
	SteamworksKey := loadEnv("SteamwroksKey")
	router := mux.NewRouter()
	dbType := "sqlite"
	core, err := core.Initialize(IGDBClientId, IGDBClientSecret, SteamworksKey, dbType)

	if err != nil {
		log.Fatal(err)
	}

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
			response := core.Client.Game.GetTrend(page)
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(response)
		}).Methods("GET")

	router.HandleFunc("/api/game/{id}",
		func(responseWriter http.ResponseWriter, request *http.Request) {
			id := strings.TrimPrefix(request.URL.Path, "/api/game/")
			response := core.GetGame(id)
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(response)
		}).Methods("GET")

	router.HandleFunc("/api/player/games/{id}",
		func(responseWriter http.ResponseWriter, request *http.Request) {
			idPlayerSteam := strings.TrimPrefix(request.URL.Path, "/api/player/games/")
			page, err := getPage(request)
			if err != nil {
				internalServerErrorHandler(responseWriter, request, err)
			}
			// response := core.Client.Game.GetPlayerGames(id, page)
			response, err := core.GetPlayerGames(idPlayerSteam, page)
			if err != nil {
				internalServerErrorHandler(responseWriter, request, err)
			}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(response)
		}).Methods("GET")

	log.Println("Server started at :8080")

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Permitir origem do Vite
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(router)))
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
