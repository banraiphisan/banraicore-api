package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/internal/domain"
	"github.com/banraiphisan/banraicore-api/internal/domain/entity"
	"github.com/banraiphisan/banraicore-api/pkg/cache"
	"github.com/banraiphisan/banraicore-api/pkg/db"
	"github.com/banraiphisan/banraicore-api/pkg/logger"
	"github.com/banraiphisan/banraicore-api/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	conf  *config.Configuration
	db    *db.DB
	Cache cache.Engine
}

func NewUserRepository(db *db.DB, _ logger.Logger, cache cache.Engine, cfg *config.Configuration) domain.UserRepository {
	return &UserRepository{
		conf:  cfg,
		db:    db,
		Cache: cache,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindByEmailOrUsername(ctx context.Context, email, username string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("email = ? OR username = ?", email, username).
		First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, err
}

func (r *UserRepository) FindRoleByCode(ctx context.Context, roleCode string) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.WithContext(ctx).Where("code = ?", roleCode).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *UserRepository) GetUsers(ctx context.Context, offset, limit int) ([]entity.User, int64, error) {
	var users []entity.User
	var total int64

	if err := r.db.WithContext(ctx).Model(&entity.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Preload("Role").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("id = ?", id).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	utils.PrintToJSON(user)
	return r.db.WithContext(ctx).Updates(user).Error
}
