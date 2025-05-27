package service

import (
	"github.com/tubfuzzy/banraicore-api/config"
	"github.com/tubfuzzy/banraicore-api/internal/domain"
	"github.com/tubfuzzy/banraicore-api/pkg/cache"
	"github.com/tubfuzzy/banraicore-api/pkg/logger"
)

type ReportService struct {
	ReportRepository domain.ReportRepository
	Cache            cache.Engine
	Logger           logger.Logger
	Conf             *config.Configuration
}

func NewReportService(reportRepository domain.ReportRepository, cache cache.Engine, logger logger.Logger, conf *config.Configuration) domain.ReportService {
	return &ReportService{
		ReportRepository: reportRepository,
		Cache:            cache,
		Logger:           logger,
		Conf:             conf,
	}
}
