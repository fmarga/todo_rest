package main

import (
	"fmt"
	"net/http"
	"todo_rest/configs"
	"todo_rest/handlers"

	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)
	r.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
