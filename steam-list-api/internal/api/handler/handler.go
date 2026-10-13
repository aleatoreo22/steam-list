package handler

import "steam-list-api.com/internal/core"

type Handler struct {
	Game   *GameHandler
	Player *PlayerHandler
}

func NewHandler(core *core.Core) *Handler {
	return &Handler{
		Game: &GameHandler{
			core: core,
		},
	}
}
