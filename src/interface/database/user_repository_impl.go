package database

import (
	"context"

	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/userdm"
	"golang.org/x/xerrors"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

var (
	u1, _ = userdm.GenForTest(userdm.NewUserID(), "user1", "lastuser1")
	u2, _ = userdm.GenForTest(userdm.NewUserID(), "user2", "lastuser2")
	u3, _ = userdm.GenForTest(userdm.NewUserID(), "user3", "lastuser3")
	users = []*userdm.User{
		u1,
		u2,
		u3,
	}
)

func (repo *UserRepositoryImpl) FindAll(ctx context.Context) ([]*userdm.User, error) {
	return users, nil
}

func (repo *UserRepositoryImpl) FindByID(ctx context.Context, id userdm.UserID) (*userdm.User, error) {
	for _, u := range users {
		if u.ID().Equal(id) {
			return u, nil
		}
	}
	return nil, xerrors.Errorf("userdm is not found: %s", id)
}
func (repo *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	// DBの保存書処理
	return nil
}
