package repository

import "ecpos/internal/model"

func NewTestRepository(r *string) UserRepository {
	return &userTestRepository{
		dummy: *r,
	}
}

type userTestRepository struct {
	dummy string
}

func (u userTestRepository) ErrorExample(id int) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userTestRepository) FirstByID(id int) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
