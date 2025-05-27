package service

import (
	"context"
	"errors"
	"github.com/segmentio/ksuid"
	"github.com/tubfuzzy/banraicore-api/config"
	"github.com/tubfuzzy/banraicore-api/internal/domain"
	"github.com/tubfuzzy/banraicore-api/internal/domain/entity"
	"github.com/tubfuzzy/banraicore-api/pkg/cache"
	"github.com/tubfuzzy/banraicore-api/pkg/logger"
	"time"
)

type ShortUrlService struct {
	ShortUrlRepository domain.ShortUrlRepository
	Cache              cache.Engine
	Logger             logger.Logger
	Conf               *config.Configuration
}

func NewShortUrlService(shortUrlRepository domain.ShortUrlRepository, cache cache.Engine, logger logger.Logger, conf *config.Configuration) domain.ShortUrlService {
	return &ShortUrlService{
		ShortUrlRepository: shortUrlRepository,
		Cache:              cache,
		Logger:             logger,
		Conf:               conf,
	}
}

func (s *ShortUrlService) Create(ctx context.Context, payload domain.CreateShortUrlPayload) (domain.ShortUrl, error) {
	code := ksuid.New().String()[:8]
	now := time.Now()

	newEntity := entity.ShortUrl{
		Code:      code,
		TargetUrl: payload.TargetUrl,
		CreatedAt: now,
		ExpiresAt: payload.ExpiresAt,
	}

	created, err := s.ShortUrlRepository.Create(ctx, newEntity)
	if err != nil {
		return domain.ShortUrl{}, err
	}

	return domain.ShortUrl{
		ID:        created.ID,
		Code:      created.Code,
		TargetUrl: created.TargetUrl,
		CreatedAt: created.CreatedAt,
		ExpiresAt: created.ExpiresAt,
	}, nil
}

func (s *ShortUrlService) GetByCode(ctx context.Context, code string) (*domain.ShortUrl, error) {
	data, err := s.ShortUrlRepository.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("not found")
	}

	return &domain.ShortUrl{
		ID:        data.ID,
		Code:      data.Code,
		TargetUrl: data.TargetUrl,
		CreatedAt: data.CreatedAt,
		ExpiresAt: data.ExpiresAt,
	}, nil
}

func (s *ShortUrlService) GetAll(ctx context.Context) ([]domain.ShortUrl, error) {
	list, err := s.ShortUrlRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []domain.ShortUrl
	for _, item := range list {
		result = append(result, domain.ShortUrl{
			ID:        item.ID,
			Code:      item.Code,
			TargetUrl: item.TargetUrl,
			CreatedAt: item.CreatedAt,
			ExpiresAt: item.ExpiresAt,
		})
	}

	return result, nil
}

func (s *ShortUrlService) DeleteByCode(ctx context.Context, code string) error {
	return s.ShortUrlRepository.DeleteByCode(ctx, code)
}
