package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World")
	router := mux.NewRouter()
	router.HandleFunc("/api/hello",
		func(responseWriter http.ResponseWriter, request *http.Request) {
			response := APIResponse{Message: "Hello, World!"}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(response)
		}).Methods("GET")
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
