package user

import (
	"unicode/utf8"

	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/id"
	"golang.org/x/xerrors"
)

type User struct {
	id        id.UserID
	firstName string
	lastName  string
}

func (u *User) ID() id.UserID {
	return u.id
}
func (u *User) FirstName() string {
	return u.firstName
}
func (u *User) LastName() string {
	return u.lastName
}

var (
	firstNameLength = 30
	lastNameLength  = 30
)

func NewUser(id id.UserID, first, last string) (*User, error) {
	if first == "" {
		return nil, xerrors.New("first name must be not empty")
	}
	if last == "" {
		return nil, xerrors.New("last name must be not empty")
	}

	if l := utf8.RuneCountInString(first); l > firstNameLength {
		return nil, xerrors.Errorf("first name must be less than %d", firstNameLength)
	}
	if l := utf8.RuneCountInString(last); l > lastNameLength {
		return nil, xerrors.Errorf("last name must be less than %d", lastNameLength)
	}

	return &User{
		id:        id,
		firstName: first,
		lastName:  last,
	}, nil
}
