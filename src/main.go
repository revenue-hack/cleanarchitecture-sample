package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/revenue-hack/cleanarchitecture-sample/src/infra/router"
)

func main() {
	userRouter, err := router.NewUserRouter()
	if err != nil {
		log.Fatal(err)
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World"})
	}))
	api.SetApp(userRouter)
	log.Fatal(http.ListenAndServe(":8000", api.MakeHandler()))
}
