package service

import (
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
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
