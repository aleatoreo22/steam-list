package routes

import (
	"github.com/gorilla/mux"
	"steam-list-api.com/internal/api/handler"
)

func PlayerRoutes(r *mux.Router, h *handler.PlayerHandler) {
	r.HandleFunc("/api/player/games/{id}", h.GetPlayerGamesById).Methods("GET")
}
