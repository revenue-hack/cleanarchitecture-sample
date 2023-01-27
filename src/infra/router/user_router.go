package router

import (
	"github.com/gin-gonic/gin"
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/controller"
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/database"
	"github.com/revenue-hack/cleanarchitecture-sample/src/interface/presenter"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
)

func NewUserRouter(g *gin.Engine) {
	g.GET("/users", func(ctx *gin.Context) {
		userRepoImpl := database.NewUserRepositoryImpl()
		err := controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).GetUserList(ctx)
		if err != nil {
			ctx.Error(err)
			return
		}
	})
	g.GET("/users/:id", func(ctx *gin.Context) {
		type reqStruct struct {
			ID string `uri:"id"`
		}
		var req reqStruct
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(400, gin.H{"status": "bad request"})
			return
		}

		in := input.GetUserByIDInput{ID: req.ID}
		userRepoImpl := database.NewUserRepositoryImpl()
		err := controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).GetUserByID(ctx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
	})
}
