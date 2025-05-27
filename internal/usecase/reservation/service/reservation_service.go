package service

import (
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/internal/domain"
	"github.com/banraiphisan/banraicore-api/pkg/cache"
	"github.com/banraiphisan/banraicore-api/pkg/logger"
)

type ReservationService struct {
	ReservationRepository domain.ReservationRepository
	Cache                 cache.Engine
	Logger                logger.Logger
	Conf                  *config.Configuration
}

func NewReservationService(reservation domain.ReservationRepository, cache cache.Engine, logger logger.Logger, conf *config.Configuration) domain.ReservationService {
	return &ReservationService{
		ReservationRepository: reservation,
		Cache:                 cache,
		Logger:                logger,
		Conf:                  conf,
	}
}
