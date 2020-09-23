package userusecase

import (
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/output"
	"golang.org/x/xerrors"
)

type GetUserByIDUsecase struct {
	userRepository UserRepository
}

func NewGetUserByID(userRepo UserRepository) *GetUserByIDUsecase {
	return &GetUserByIDUsecase{
		userRepository: userRepo,
	}
}

func (use *GetUserByIDUsecase) Exec(in *input.GetUserByIDInput) (*output.UserByID, error) {
	if in.ID == "" {
		return nil, xerrors.New("id must be not empty")
	}
	user, err := use.userRepository.FindByID(in.ID)
	if err != nil {
		return nil, err
	}
	return &output.UserByID{
		ID:        user.ID(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
	}, nil
}
