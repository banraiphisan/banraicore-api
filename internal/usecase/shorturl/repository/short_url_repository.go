package repository

import (
	"context"
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain/entity"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/db"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
	"gorm.io/gorm"
)

type ShortUrlRepository struct {
	conf  *config.Configuration
	db    *db.DB
	Cache cache.Engine
}

func NewShortUrlRepository(db *db.DB, _ logger.Logger, cache cache.Engine, cfg *config.Configuration) domain.ShortUrlRepository {
	return &ShortUrlRepository{
		conf:  cfg,
		db:    db,
		Cache: cache,
	}
}

func (r *ShortUrlRepository) Create(ctx context.Context, shortUrl entity.ShortUrl) (entity.ShortUrl, error) {
	err := r.db.WithContext(ctx).Create(&shortUrl).Error
	if err != nil {
		return entity.ShortUrl{}, err
	}
	return shortUrl, nil
}

func (r *ShortUrlRepository) GetByCode(ctx context.Context, code string) (*entity.ShortUrl, error) {
	var shortUrl entity.ShortUrl
	err := r.db.WithContext(ctx).Where("short_code = ?", code).First(&shortUrl).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &shortUrl, nil
}

func (r *ShortUrlRepository) GetAll(ctx context.Context) ([]entity.ShortUrl, error) {
	var list []entity.ShortUrl
	err := r.db.WithContext(ctx).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ShortUrlRepository) DeleteByCode(ctx context.Context, code string) error {
	return r.db.WithContext(ctx).Where("short_code = ?", code).Delete(&entity.ShortUrl{}).Error
}
