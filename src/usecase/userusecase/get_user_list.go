package userusecase

import (
	"context"

	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/user"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/output"
)

type GetUserListUsecase struct {
	userRepository user.UserRepository
}

func NewGetUserList(userRepo user.UserRepository) *GetUserListUsecase {
	return &GetUserListUsecase{
		userRepository: userRepo,
	}
}

func (use *GetUserListUsecase) Exec(ctx context.Context) (*output.UserList, error) {
	users, err := use.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	userItems := make([]*output.UserItem, len(users))
	for i, user := range users {
		userItems[i] = &output.UserItem{
			ID:        user.ID().String(),
			FirstName: user.FirstName(),
			LastName:  user.LastName(),
		}
	}
	return &output.UserList{Users: userItems}, nil
}
