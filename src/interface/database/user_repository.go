package database

import (
	"github.com/revenue-hack/cleanarchitecture-sample/src/entity"
	"golang.org/x/xerrors"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

var (
	u1, _ = entity.NewUser("1", "user1", "lastuser1")
	u2, _ = entity.NewUser("2", "user2", "lastuser2")
	u3, _ = entity.NewUser("3", "user3", "lastuser3")
	users = []*entity.User{
		u1,
		u2,
		u3,
	}
)

func (repo *UserRepositoryImpl) FindAll() ([]*entity.User, error) {
	return users, nil
}

func (repo *UserRepositoryImpl) FindByID(id string) (*entity.User, error) {
	if id == "" {
		return nil, xerrors.New("id must be not empty")
	}
	for _, user := range users {
		if user.ID() == id {
			return user, nil
		}
	}
	return nil, xerrors.Errorf("user is not found: %s", id)
}
