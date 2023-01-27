//go:generate mockgen -source ./user_repository.go -destination ../../mock/mock_user/user_repository_mock.go
package user

import (
	"context"

	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/id"
)

type UserRepository interface {
	//Store(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, userID id.UserID) (*User, error)
}
