package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/banraiphisan/banraicore-api/internal/domain/entity"
)

type AuthService interface {
	// TODO: implement this
	//Authentication(ctx context.Context, username, password string) (string, string, error)
	//RefreshToken(ctx context.Context, accessToken, refreshToken string) (string, string, error)
	//Logout(ctx context.Context, accessToken, refreshToken string) error
	GetUserRolePermissions(ctx context.Context, userID string) ([]RolePermission, error)
}

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserRolePermissions(ctx context.Context, userID string) ([]entity.RolePermission, error)
}

// DTO for returning role permissions
type RolePermission struct {
	RoleID       uuid.UUID `json:"roleId"`
	PermissionID uuid.UUID `json:"permissionId"`
	Permission   string    `json:"permission"`
}
