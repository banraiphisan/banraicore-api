package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/internal/domain"
)

type RoomHandler struct {
	config.Configuration
	domain.RoomService
}

func NewRoomHandler(roomService domain.RoomService, config *config.Configuration) RoomHandler {
	return RoomHandler{
		RoomService:   roomService,
		Configuration: *config,
	}
}

func (h RoomHandler) InitRoute(app fiber.Router) {

}
