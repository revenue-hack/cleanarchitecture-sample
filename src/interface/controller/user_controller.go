package controller

import (
	"context"

	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/user"
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/presenter"
	"github.com/revenue-hack/cleanarchitecture-sample/src/support/smperr"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
)

type userController struct {
	delivery presenter.UserPresenter
	userRepo user.UserRepository
}

func NewUserController(p presenter.UserPresenter, userRepo user.UserRepository) *userController {
	return &userController{
		delivery: p,
		userRepo: userRepo,
	}
}

func (c *userController) GetUserList(ctx context.Context) error {
	usecase := userusecase.NewGetUserList(c.userRepo)
	out, err := usecase.Exec(ctx)
	if err != nil {
		return err
	}
	return smperr.BadRequestf("ああああ: %+v", 1)
	c.delivery.UserList(out)
	return nil

}

func (c *userController) GetUserByID(ctx context.Context, in *input.GetUserByIDInput) error {
	usecase := userusecase.NewGetUserByID(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.UserByID(out)
	return nil
}
