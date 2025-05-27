package domain

import (
	"context"
	"github.com/banraiphisan/banraicore-api/internal/domain/entity"
	"time"
)

type ShortUrlService interface {
	Create(ctx context.Context, payload CreateShortUrlPayload) (ShortUrl, error)
	GetByCode(ctx context.Context, code string) (*ShortUrl, error)
	GetAll(ctx context.Context) ([]ShortUrl, error)
	DeleteByCode(ctx context.Context, code string) error
}

type ShortUrlRepository interface {
	Create(ctx context.Context, shortUrl entity.ShortUrl) (entity.ShortUrl, error)
	GetByCode(ctx context.Context, code string) (*entity.ShortUrl, error)
	GetAll(ctx context.Context) ([]entity.ShortUrl, error)
	DeleteByCode(ctx context.Context, code string) error
}

type CreateShortUrlPayload struct {
	TargetUrl string     `json:"target_url"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

type ShortUrl struct {
	ID        int64      `json:"id"`
	Code      string     `json:"code"`
	TargetUrl string     `json:"target_url"`
	CreatedAt time.Time  `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}
