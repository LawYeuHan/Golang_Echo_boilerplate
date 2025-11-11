package service

import (
	"ecpos/internal/model"
	"ecpos/internal/repository"
)

type UserService interface {
	GetUserByID(id int) (*model.User, error)
	GetUserByIDWithError(id int) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{ur}
}

func (us *userService) GetUserByID(id int) (*model.User, error) {
	return us.userRepository.FirstByID(id)
}

func (us *userService) GetUserByIDWithError(id int) (*model.User, error) {
	return us.userRepository.ErrorExample(id)
}
