package repository

import (
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/internal/domain"
	"github.com/banraiphisan/banraicore-api/pkg/cache"
	"github.com/banraiphisan/banraicore-api/pkg/db"
	"github.com/banraiphisan/banraicore-api/pkg/logger"
)

type RoomRepository struct {
	conf  *config.Configuration
	db    *db.DB
	Cache cache.Engine
}

func NewRoomRepository(db *db.DB, _ logger.Logger, cache cache.Engine, cfg *config.Configuration) domain.RoomRepository {
	return &RoomRepository{
		conf:  cfg,
		db:    db,
		Cache: cache,
	}
}
