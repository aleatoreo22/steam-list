package routes

import (
	"github.com/gorilla/mux"
	"steam-list-api.com/internal/api/handler"
)

func GameRoutes(r *mux.Router, h *handler.GameHandler) {
	r.HandleFunc("/api/game/trend", h.GetGameTrend).Methods("GET")
	r.HandleFunc("/api/game/{id}", h.GetGameById).Methods("GET")
}
