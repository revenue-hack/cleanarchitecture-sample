package userusecase

import (
	"context"

	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/userdm"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/userinput"
)

type UpdateUserUsecase struct {
	userRepository userdm.UserRepository
}

func NewUpdateUser(userRepo userdm.UserRepository) *UpdateUserUsecase {
	return &UpdateUserUsecase{
		userRepository: userRepo,
	}
}

func (use *UpdateUserUsecase) Exec(ctx context.Context, userID userdm.UserID, in *userinput.UpdateUserInput) error {
	user, err := use.userRepository.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if err := user.Update(in.FirstName, in.LastName); err != nil {
		return err
	}

	if err := use.userRepository.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
