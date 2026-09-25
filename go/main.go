package main

import(
	"net/http"
	"errors"
	"fmt"
	"math/rand"

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
	r.Get("/", recipeHandler.ListRecipes)
	r.Post("/", recipeHandler.CreateRecipe)
	r.Route("/{recipeId}", func(r chi.Router) {
		r.Use(recipeHandler.RecipeCtx)
		r.Get("/", recipeHandler.GetRecipe)
		r.Put("/", recipeHandler.UpdateRecipe)
		r.Delete("/", recipeHandler.DeleteRecipe)
	})

	return r
}

func dbGetRecipe(recipeId string) (*Recipe, error) {
	for _, r := range(recipesDb) {
		if r.ID == recipeId {
			return r, nil
		}
	}
	return nil, errors.New("recipe not found")
}

func dbNewRecipe (recipe *Recipe) (string, error) {
	recipe.ID = fmt.Sprintf("%d", rand.Intn(100) + 10)
	recipesDb = append(recipesDb, recipe)
	return recipe.ID, nil
}

func dbUpdateRecipe(id string, recipe *Recipe) (*Recipe, error) {
	for i, r := range(recipesDb) {
		if r.ID == id {
			recipesDb[i] = recipe
			return recipe, nil
		}
	}
	return nil, errors.New("recipe not found")
}
