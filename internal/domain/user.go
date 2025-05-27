package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/banraiphisan/banraicore-api/internal/domain/entity"
)

type UserService interface {
	// TODO: implement this
	CreateUser(ctx context.Context, createUser CreateOrUpdateUserPayload) error
	UpdateUser(ctx context.Context, createUser CreateOrUpdateUserPayload) error
	GetUsers(ctx context.Context, page, pageSize int) ([]UserInfo, int64, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (UserInfo, error)
	//DeleteUserByID(ctx context.Context, id string) error
}

type UserRepository interface {
	// TODO: implement this
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUsers(ctx context.Context, offset, limit int) ([]entity.User, int64, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	//DeleteUserByID(ctx context.Context, id uuid.UUID) error
	FindByEmailOrUsername(ctx context.Context, email, username string) (*entity.User, error)
	FindRoleByCode(ctx context.Context, roleCode string) (*entity.Role, error)
}

// Payload used for creating or updating a user
type CreateOrUpdateUserPayload struct {
	ID       *uuid.UUID `json:"id,omitempty"`
	Username string     `json:"username" validate:"required,min=3,max=50"`
	Email    string     `json:"email" validate:"required,email"`
	Password string     `json:"password,omitempty" validate:"omitempty,min=6,required_without=ID"`
	RoleCode string     `json:"roleCode" validate:"required,oneof=ADMIN RECEPTIONIST HOUSEKEEPING"`
}

// DTO for returning user information without exposing sensitive fields
type UserInfo struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	RoleID   uuid.UUID `json:"roleId"`
	RoleName string    `json:"roleName"`
}
