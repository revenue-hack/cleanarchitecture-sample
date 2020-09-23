package controller

import (
	"github.com/revenue-hack/cleanarchitecture-sample/src/entity"
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/database"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
)

type UserController struct{}

func (c *UserController) GetUserList() ([]*entity.User, error) {
	userRepo := database.NewUserRepositoryImpl()
	usecase := userusecase.NewGetUserList(userRepo)
	return usecase.Exec()
}

func (c *UserController) GetUserByID(in *input.GetUserByIDInput) (*entity.User, error) {
	userRepo := database.NewUserRepositoryImpl()
	usecase := userusecase.NewGetUserByID(userRepo)
	return usecase.Exec(in)
}
