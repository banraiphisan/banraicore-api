package repository

import (
	"context"
	"errors"
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain/entity"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/db"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
	"gorm.io/gorm"
)

type AuthRepository struct {
	conf  *config.Configuration
	db    *db.DB
	Cache cache.Engine
}

func NewAuthRepository(db *db.DB, _ logger.Logger, cache cache.Engine, cfg *config.Configuration) domain.AuthRepository {
	return &AuthRepository{
		conf:  cfg,
		db:    db,
		Cache: cache,
	}
}

// GetUserByUsername retrieves a user by username
func (r *AuthRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	result := r.db.WithContext(ctx).Where("username = ?", username).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// GetUserRolePermissions retrieves role permissions for a user
func (r *AuthRepository) GetUserRolePermissions(ctx context.Context, userID string) ([]entity.RolePermission, error) {
	var rolePermissions []entity.RolePermission

	result := r.db.WithContext(ctx).
		Joins("JOIN users ON users.role_id = roles.id").
		Joins("JOIN role_permissions ON role_permissions.role_id = roles.id").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Preload("Permissions").
		Preload("Role").
		Where("users.id = ?", userID).
		Find(&rolePermissions)

	if result.Error != nil {
		return nil, result.Error
	}

	return rolePermissions, nil
}
