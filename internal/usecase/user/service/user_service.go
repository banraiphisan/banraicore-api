package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/internal/domain"
	"github.com/banraiphisan/banraicore-api/internal/domain/entity"
	"github.com/banraiphisan/banraicore-api/pkg/cache"
	"github.com/banraiphisan/banraicore-api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository domain.UserRepository
	Cache          cache.Engine
	Logger         logger.Logger
	Conf           *config.Configuration
}

func NewUserService(userRepository domain.UserRepository, cache cache.Engine, logger logger.Logger, conf *config.Configuration) domain.UserService {
	return &UserService{
		UserRepository: userRepository,
		Cache:          cache,
		Logger:         logger,
		Conf:           conf,
	}
}

func (s UserService) CreateUser(ctx context.Context, createUser domain.CreateOrUpdateUserPayload) error {
	existingUser, err := s.UserRepository.FindByEmailOrUsername(ctx, createUser.Email, createUser.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return fmt.Errorf("user already exists")
	}

	role, err := s.UserRepository.FindRoleByCode(ctx, createUser.RoleCode)
	if err != nil {
		return fmt.Errorf("invalid role code: %s", createUser.RoleCode)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		ID:           uuid.New(),
		Username:     createUser.Username,
		Email:        createUser.Email,
		PasswordHash: string(hashedPassword),
		RoleID:       &role.ID,
	}

	if err := s.UserRepository.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserService) GetUsers(ctx context.Context, page, pageSize int) ([]domain.UserInfo, int64, error) {
	offset := (page - 1) * pageSize
	users, total, err := s.UserRepository.GetUsers(ctx, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var userInfos []domain.UserInfo
	for _, user := range users {
		userInfos = append(userInfos, domain.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			RoleID:   *user.RoleID,
			RoleName: user.Role.Name,
		})
	}

	return userInfos, total, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (domain.UserInfo, error) {
	user, err := s.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		return domain.UserInfo{}, fmt.Errorf("user not found")
	}

	return domain.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		RoleID:   *user.RoleID,
		RoleName: user.Role.Name,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, updateUser domain.CreateOrUpdateUserPayload) error {
	userID := *updateUser.ID
	existingUser, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	role, err := s.UserRepository.FindRoleByCode(ctx, updateUser.RoleCode)
	if err != nil {
		return fmt.Errorf("invalid role code: %s", updateUser.RoleCode)
	}

	existingUser.RoleID = &role.ID
	existingUser.Role = role
	
	err = s.UserRepository.UpdateUser(ctx, existingUser)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
