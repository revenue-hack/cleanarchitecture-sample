package controller

import (
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/database"
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/presenter"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
)

type userController struct {
	delivery presenter.UserPresenter
	userRepo *database.UserRepositoryImpl
}

func NewUserController(p presenter.UserPresenter, userRepo *database.UserRepositoryImpl) *userController {
	return &userController{
		delivery: p,
		userRepo: userRepo,
	}
}

func (c *userController) GetUserList() error {
	usecase := userusecase.NewGetUserList(c.userRepo)
	out, err := usecase.Exec()
	if err != nil {
		return err
	}
	return c.delivery.UserList(out)

}

func (c *userController) GetUserByID(in *input.GetUserByIDInput) error {
	usecase := userusecase.NewGetUserByID(c.userRepo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}
	return c.delivery.UserByID(out)
}
