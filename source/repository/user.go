package repository

import (
	"basic/source/model"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() (*[]model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}
type userRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

func (r *userRepository) GetUsers() (*[]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {

		return nil, errors.Wrap(err, "failed to get all users")
	}

	return &users, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {

		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return &user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}

	return &user, nil
}
