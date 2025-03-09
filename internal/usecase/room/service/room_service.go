package service

import (
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
)

type RoomService struct {
	RoomRepository domain.RoomRepository
	Cache          cache.Engine
	Logger         logger.Logger
	Conf           *config.Configuration
}

func NewRoomService(roomRepository domain.RoomRepository, cache cache.Engine, logger logger.Logger, conf *config.Configuration) domain.RoomService {
	return &RoomService{
		RoomRepository: roomRepository,
		Cache:          cache,
		Logger:         logger,
		Conf:           conf,
	}
}
