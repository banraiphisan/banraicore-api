package service

import (
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/internal/domain"
	"github.com/banraiphisan/banraicore-api/pkg/cache"
	"github.com/banraiphisan/banraicore-api/pkg/logger"
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
