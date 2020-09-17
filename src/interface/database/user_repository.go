package database

import (
	"github.com/revenue-hack/cleanarchitecture-sample/src/entity"
	"golang.org/x/xerrors"
)

type UserRepositoryImpl struct{}

func NewUserrepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

var (
	users = []*entity.User{
		{
			ID:        "1",
			FirstName: "user1",
			LastName:  "lastuser1",
		},
		{
			ID:        "2",
			FirstName: "user2",
			LastName:  "lastuser2",
		},
		{
			ID:        "3",
			FirstName: "user3",
			LastName:  "lastuser3",
		},
	}
)

func (repo *UserRepositoryImpl) FindAll() ([]*entity.User, error) {
	return users, nil
}

func (repo *UserRepositoryImpl) FindByID(id string) (*entity.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, xerrors.Errorf("user is not found: %s", id)
}
