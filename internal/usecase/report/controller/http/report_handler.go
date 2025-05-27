package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tubfuzzy/banraicore-api/config"
	"github.com/tubfuzzy/banraicore-api/internal/domain"
)

type ReportHandler struct {
	config.Configuration
	domain.ReportService
}

func NewReportHandler(reportService domain.ReportService, config *config.Configuration) ReportHandler {
	return ReportHandler{
		ReportService: reportService,
		Configuration: *config,
	}
}

func (h ReportHandler) InitRoute(app fiber.Router) {

}
