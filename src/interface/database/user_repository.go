package database

import (
	"context"

	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/id"
	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/user"
	"golang.org/x/xerrors"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

var (
	u1, _ = user.NewUser("1", "user1", "lastuser1")
	u2, _ = user.NewUser("2", "user2", "lastuser2")
	u3, _ = user.NewUser("3", "user3", "lastuser3")
	users = []*user.User{
		u1,
		u2,
		u3,
	}
)

func (repo *UserRepositoryImpl) FindAll(ctx context.Context) ([]*user.User, error) {
	return users, nil
}

func (repo *UserRepositoryImpl) FindByID(ctx context.Context, id id.UserID) (*user.User, error) {

	for _, u := range users {
		if u.ID().Equal(id) {
			return u, nil
		}
	}
	return nil, xerrors.Errorf("user is not found: %s", id)
}
