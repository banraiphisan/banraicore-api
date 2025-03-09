package repository

import (
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/db"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
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
