package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain"
)

type ReservationHandler struct {
	config.Configuration
	domain.ReservationService
}

func NewReservationHandler(reportService domain.ReservationService, config *config.Configuration) ReservationHandler {
	return ReservationHandler{
		ReservationService: reportService,
		Configuration:      *config,
	}
}

func (h ReservationHandler) InitRoute(app fiber.Router) {

}
