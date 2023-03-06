//go:generate mockgen -source ./user_repository.go -destination ../../mock/mock_user/user_repository_mock.go
package userdm

import (
	"context"
)

type UserRepository interface {
	Store(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, userID UserID) (*User, error)
}
