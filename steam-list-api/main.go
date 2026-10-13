package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"steam-list-api.com/internal/api/handler"
	"steam-list-api.com/internal/api/routes"
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

	handler := handler.NewHandler(core)
	routes.GameRoutes(router, handler.Game)
	routes.PlayerRoutes(router, handler.Player)

	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/api/hello",
		func(responseWriter http.ResponseWriter, request *http.Request) {
			response := model.APIResponse{Message: "Hello, World!"}
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
