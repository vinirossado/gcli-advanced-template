package service

import (
	"basic/source/model"
	"basic/source/repository"
)

type UserService interface {
	GetUsers() ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
	CreateUser(*model.User) (int, error)
	UpdateUser(*model.User) (bool, error)
	DeleteUser(id int) (bool, error)
}

type userService struct {
	*Service
	userRepository repository.UserRepository
}

func NewUserService(service *Service, userRepository repository.UserRepository) UserService {
	return &userService{
		Service:        service,
		userRepository: userRepository,
	}
}

func (s *userService) GetUsers() ([]*model.User, error) {
	return s.userRepository.GetUsers()
}

func (s *userService) GetUserById(id int) (*model.User, error) {
	return s.userRepository.GetUserById(id)
}

func (s *userService) CreateUser(user *model.User) (int, error) {
	return s.userRepository.CreateUser(user)
}

func (s *userService) UpdateUser(user *model.User) (bool, error) {
	return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUser(id int) (bool, error) {
	return s.userRepository.DeleteUser(id)
}
