package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/revenue-hack/cleanarchitecture-sample/src/infra/middleware"
	"github.com/revenue-hack/cleanarchitecture-sample/src/infra/router"
)

func main() {
	g := gin.Default()

	g.Use(middleware.HandleErrorMiddleware())

	router.NewUserRouter(g)

	log.Fatal(g.Run(":8000"))
}
