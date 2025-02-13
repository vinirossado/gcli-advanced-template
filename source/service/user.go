package service

import (
	"basic/pkg/cache"
	"basic/pkg/helper/mapper"
	"basic/pkg/helper/uuid"
	"basic/source/model"
	"basic/source/repository"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email" binding:"required,email"`
	Avatar   string `json:"avatar"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type UserService interface {
	Register(ctx context.Context, req *RegisterRequest) error
	Login(ctx context.Context, req *LoginRequest) (string, error)
	GetProfile(ctx context.Context, userID string) (*UserResponse, error)
	UpdateProfile(ctx context.Context, userID string, req *UpdateProfileRequest) error
	GenerateToken(ctx context.Context, userID string) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func NewUserService(service *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

func (s *userService) Register(ctx context.Context, req *RegisterRequest) error {
	if user, err := s.userRepo.GetByUsername(ctx, req.Username); err == nil && user != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}

	user := &model.User{
		UserID:   uuid.GenUUID(),
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}
	if err = s.userRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func (s *userService) Login(ctx context.Context, req *LoginRequest) (string, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return "", errors.Wrap(err, "failed to get user by username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.Wrap(err, "failed to hash password")
	}
	token, err := s.GenerateToken(ctx, user.UserID)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}
	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userID string) (*UserResponse, error) {
	if cache.Cache == nil {
		fmt.Println("Cache not initialized.")
		return &UserResponse{}, nil
	}

	if entry, err := cache.Cache.Get("user"); err == nil {
		var userResponse UserResponse
		err = json.Unmarshal(entry, &userResponse)
		if err == nil {
			return &userResponse, nil
		}
	}

	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	var userResponse UserResponse
	mapper.Map(user, &userResponse)

	responseBytes, err := json.Marshal(userResponse)
	if err == nil {
		err = cache.Cache.Set("user", responseBytes)
		if err != nil {
			fmt.Println("Fail to save cache:", err)
		}
	}
	return &userResponse, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userID string, req *UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.Wrap(err, "failed to get user by ID")
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.userRepo.Update(ctx, user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (s *userService) GenerateToken(ctx context.Context, userID string) (string, error) {
	token, err := s.jwt.GenToken(userID, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}

	return token, nil
}
