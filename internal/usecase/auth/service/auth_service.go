package service

import (
	"context"
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/internal/domain"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
)

type AuthService struct {
	AuthRepository domain.AuthRepository
	Cache          cache.Engine
	Logger         logger.Logger
	Conf           *config.Configuration
}

func NewAuthService(authRepository domain.AuthRepository, cache cache.Engine, logger logger.Logger, conf *config.Configuration) domain.AuthService {
	return &AuthService{
		AuthRepository: authRepository,
		Cache:          cache,
		Logger:         logger,
		Conf:           conf,
	}
}

func (s AuthService) GetUserRolePermissions(ctx context.Context, userID string) ([]domain.RolePermission, error) {
	results := make([]domain.RolePermission, 0)
	rolePermissions, err := s.AuthRepository.GetUserRolePermissions(ctx, userID)
	if err != nil {
		return results, err
	}

	for _, rolePermission := range rolePermissions {
		results = append(results, domain.RolePermission{
			RoleID:       rolePermission.RoleID,
			PermissionID: rolePermission.PermissionID,
			Permission:   rolePermission.Permission.Name,
		})
	}

	return results, nil
}
