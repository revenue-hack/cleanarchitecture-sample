package router

import (
	"context"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/controller"
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/database"
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/presenter"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
)

func NewUserRouter() (rest.App, error) {
	router, err := rest.MakeRouter(
		rest.Get("/users", func(w rest.ResponseWriter, req *rest.Request) {
			ctx := context.Background()
			userRepoImpl := database.NewUserRepositoryImpl()
			err := controller.NewUserController(presenter.NewUserPresenter(w), userRepoImpl).GetUserList(ctx)
			if err != nil {
				rest.Error(w, err.Error(), 500)
				return
			}
		}),
		rest.Get("/users/#id", func(w rest.ResponseWriter, req *rest.Request) {
			id := req.PathParam("id")
			in := input.GetUserByIDInput{ID: id}
			ctx := context.Background()
			userRepoImpl := database.NewUserRepositoryImpl()
			err := controller.NewUserController(presenter.NewUserPresenter(w), userRepoImpl).GetUserByID(ctx, &in)
			if err != nil {
				rest.Error(w, err.Error(), 500)
				return
			}
		}),
	)
	if err != nil {
		return nil, err
	}
	return router, nil
}
