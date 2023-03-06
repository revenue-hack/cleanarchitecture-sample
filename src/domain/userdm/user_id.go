package userdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}
func NewUserIDByVal(val string) (UserID, error) {
	if val == "" {
		return UserID(""), xerrors.New("userdm id must not be empty")
	}
	return UserID(val), nil
}

func (id UserID) String() string {
	return string(id)
}

func (id UserID) Equal(userID2 UserID) bool {
	return string(id) == string(userID2)
}
