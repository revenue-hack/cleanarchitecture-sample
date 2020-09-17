package userusecase

import (
	"github.com/revenue-hack/cleanarchitecture-sample/src/entity"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
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

func (use *GetUserByIDUsecase) Exec(in *input.GetUserByIDInput) (*entity.User, error) {
	if in.ID == "" {
		return nil, xerrors.New("id must be not empty")
	}
	return use.userRepository.FindByID(in.ID)
}
