package userusecase

import "github.com/revenue-hack/cleanarchitecture-sample/src/entity"

type UserRepository interface {
	FindAll() ([]*entity.User, error)
	FindByID(id string) (*entity.User, error)
}
