package service

import (
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
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
