package main

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lib/pq"
)

type recipe struct {
	Name         string
	Ingredients  []string
	Instructions []string
	CookingTime  int64
	PeopleCount  int64
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres_db:5432/recipes")
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":3000", router)
}

func getRecipes(w http.ResponseWriter, r *http.Request) {

}
