package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"steam-list-api.com/internal/api/apihelper"
	"steam-list-api.com/internal/core"
)

type PlayerHandler struct {
	handler *Handler
	core    *core.Core
}

func NewPlayerHandler(handler *Handler, core *core.Core) *PlayerHandler {
	return &PlayerHandler{
		handler: handler,
		core:    core,
	}
}

func (h *PlayerHandler) GetPlayerGamesById(responseWriter http.ResponseWriter, request *http.Request) {
	idPlayerSteam := strings.TrimPrefix(request.URL.Path, "/api/player/games/")
	page, err := apihelper.GetPage(request)
	if err != nil {
		apihelper.InternalServerErrorHandler(responseWriter, request, err)
	}
	response, err := h.core.Service.Game.GetSteamPlayerGames(idPlayerSteam, page)

	if err != nil {
		apihelper.InternalServerErrorHandler(responseWriter, request, err)
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(response)
}
