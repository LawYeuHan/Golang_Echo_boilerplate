package repository

import (
	"ecpos/internal/model"
	"ecpos/pkg/log"
	"errors"
)

type UserRepository interface {
	FirstByID(id int) (*model.User, error)
	ErrorExample(id int) (*model.User, error)
}

type userRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

func (r *userRepository) FirstByID(id int) (*model.User, error) {
	user := &model.User{
		ID:       1,
		Username: "test",
		Email:    "test@example.com",
	}
	return user, nil
}

func (r *userRepository) ErrorExample(id int) (*model.User, error) {

	return nil, log.CustomError(log.ErrCodeRepo, errors.New("something went wrong"))
}
