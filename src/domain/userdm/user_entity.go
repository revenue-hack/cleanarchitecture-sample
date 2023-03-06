package userdm

import (
	"unicode/utf8"

	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/shared"
	"golang.org/x/xerrors"
)

type User struct {
	id        UserID
	firstName string
	lastName  string
	createdAt shared.CreatedAt
}

func (u *User) ID() UserID {
	return u.id
}
func (u *User) FirstName() string {
	return u.firstName
}
func (u *User) LastName() string {
	return u.lastName
}
func (u *User) CreatedAt() shared.CreatedAt {
	return u.createdAt
}

var (
	firstNameLength = 30
	lastNameLength  = 30
)

func newUser(id UserID, first, last string, createdAt shared.CreatedAt) (*User, error) {
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
		createdAt: createdAt,
	}, nil
}
