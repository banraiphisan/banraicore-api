package repository

import (
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/db"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
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
