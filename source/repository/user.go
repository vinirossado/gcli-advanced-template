package repository

import "basic/source/model"

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
	CreateUser(*model.User) (int, error)
	UpdateUser(*model.User) (bool, error)
	DeleteUser(id int) (bool, error)
}
type userRepository struct {
	*Repository
}

func NewUserRepository(repository *Repository) UserRepository {
	return &userRepository{
		Repository: repository,
	}
}

func (r *userRepository) GetUsers() ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) GetUserById(id int) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) CreateUser(user *model.User) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) UpdateUser(user *model.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) DeleteUser(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}
