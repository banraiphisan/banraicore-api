package repository

import (
	"github.com/tubfuzzy/banraicore-api/config"
	"github.com/tubfuzzy/banraicore-api/internal/domain"
	"github.com/tubfuzzy/banraicore-api/pkg/cache"
	"github.com/tubfuzzy/banraicore-api/pkg/db"
	"github.com/tubfuzzy/banraicore-api/pkg/logger"
)

type ReservationRepository struct {
	conf  *config.Configuration
	db    *db.DB
	Cache cache.Engine
}

func NewReservationRepository(db *db.DB, _ logger.Logger, cache cache.Engine, cfg *config.Configuration) domain.ReservationRepository {
	return &ReservationRepository{
		conf:  cfg,
		db:    db,
		Cache: cache,
	}
}
