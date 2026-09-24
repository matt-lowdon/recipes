package main

import(
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	router.Mount("/recipes", RecipeRoutes())
	
	http.ListenAndServe(":3000", router)
}

func RecipeRoutes() chi.Router {
	r := chi.NewRouter()
	recipeHandler := RecipeHandler{}
	r.get("/", recipeHandler.ListRecipes)
	r.Mount
