package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"steam-list-api.com/internal/api/apihelper"
	"steam-list-api.com/internal/core"
)

type GameHandler struct {
	handler *Handler
	core    *core.Core
}

func NewGameHandler(handler *Handler, core *core.Core) *GameHandler {
	return &GameHandler{
		handler: handler,
		core:    core,
	}
}

func (h *GameHandler) GetGameById(responseWriter http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, "/api/game/")
	response := h.core.Service.Game.GetGameByID(id)
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(response)
}

func (h *GameHandler) GetGameTrend(responseWriter http.ResponseWriter, request *http.Request) {
	page, err := apihelper.GetPage(request)
	if err != nil {
		apihelper.InternalServerErrorHandler(responseWriter, request, err)
	}
	response := h.core.Service.Game.GetTrend(page)
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(response)
}
