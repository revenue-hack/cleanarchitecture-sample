package userusecase

import (
	"context"

	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/id"
	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/user"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/output"
)

type GetUserByIDUsecase struct {
	userRepository user.UserRepository
}

func NewGetUserByID(userRepo user.UserRepository) *GetUserByIDUsecase {
	return &GetUserByIDUsecase{
		userRepository: userRepo,
	}
}

func (use *GetUserByIDUsecase) Exec(ctx context.Context, in *input.GetUserByIDInput) (*output.UserByID, error) {
	userID, err := id.NewUserIDByVal(in.ID)
	if err != nil {
		return nil, err
	}
	user, err := use.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &output.UserByID{
		ID:        user.ID().String(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
	}, nil
}
