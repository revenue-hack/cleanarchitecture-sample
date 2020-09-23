package userusecase

import (
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/output"
)

type GetUserListUsecase struct {
	userRepository UserRepository
}

func NewGetUserList(userRepo UserRepository) *GetUserListUsecase {
	return &GetUserListUsecase{
		userRepository: userRepo,
	}
}

func (use *GetUserListUsecase) Exec() (*output.UserList, error) {
	users, err := use.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	userItems := make([]*output.UserItem, len(users))
	for i, user := range users {
		userItems[i] = &output.UserItem{
			ID:        user.ID(),
			FirstName: user.FirstName(),
			LastName:  user.LastName(),
		}
	}
	return &output.UserList{Users: userItems}, nil
}
