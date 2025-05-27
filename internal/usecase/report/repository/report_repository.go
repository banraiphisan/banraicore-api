package repository

import (
	"github.com/tubfuzzy/banraicore-api/config"
	"github.com/tubfuzzy/banraicore-api/internal/domain"
	"github.com/tubfuzzy/banraicore-api/pkg/cache"
	"github.com/tubfuzzy/banraicore-api/pkg/db"
	"github.com/tubfuzzy/banraicore-api/pkg/logger"
)

type ReportRepository struct {
	conf  *config.Configuration
	db    *db.DB
	Cache cache.Engine
}

func NewReportRepository(db *db.DB, _ logger.Logger, cache cache.Engine, cfg *config.Configuration) domain.ReportRepository {
	return &ReportRepository{
		conf:  cfg,
		db:    db,
		Cache: cache,
	}
}
