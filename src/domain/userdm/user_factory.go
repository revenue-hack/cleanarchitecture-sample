package userdm

import (
	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/shared"
)

func GenWhenCreate(first, last string) (*User, error) {
	return newUser(NewUserID(), first, last, shared.NewCreatedAt())
}
func GenForTest(id UserID, first, last string) (*User, error) {
	return newUser(id, first, last, shared.NewCreatedAt())
}
