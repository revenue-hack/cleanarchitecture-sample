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
	updatedAt shared.UpdatedAt
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

func validateFirstName(first string) error {
	if first == "" {
		return xerrors.New("first name must be not empty")
	}
	if l := utf8.RuneCountInString(first); l > firstNameLength {
		return xerrors.Errorf("first name must be less than %d", firstNameLength)
	}

	return nil
}

func validateLastName(last string) error {
	if last == "" {
		return xerrors.New("last name must be not empty")
	}

	if l := utf8.RuneCountInString(last); l > lastNameLength {
		return xerrors.Errorf("last name must be less than %d", lastNameLength)
	}

	return nil
}

func newUser(id UserID, first, last string, createdAt shared.CreatedAt) (*User, error) {
	if err := validateFirstName(first); err != nil {
		return nil, err
	}
	if err := validateLastName(last); err != nil {
		return nil, err
	}

	return &User{
		id:        id,
		firstName: first,
		lastName:  last,
		createdAt: createdAt,
	}, nil
}

func (u *User) Update(first, last string) error {
	if err := validateFirstName(first); err != nil {
		return err
	}
	if err := validateLastName(last); err != nil {
		return err
	}

	u.firstName = first
	u.lastName = last
	u.updatedAt = shared.NewUpdatedAt()

	return nil
}
