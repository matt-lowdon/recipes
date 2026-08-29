package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type recipe struct {
	Name string
	Ingredients  []string
	Instructions []string
	CookingTime  int64
	PeopleCount  int64
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":3000", router)
}

func getRecipes
