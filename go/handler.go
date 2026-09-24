package main

import(
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RecipeHandler struct{
}

func (rh *RecipeHandler) ListRecipes(w http.ResponseWriter, *http.Request) {}
func (rh *RecipeHandler) GetRecipe(w http.ResponseWriter, *http.Request) {}
func (rh *RecipeHandler) CreateRecipe(w http.ResponseWriter, *http.Request) {}
func (rh *RecipeHandler) UpdateRecipe(w http.ResponseWriter, *http.Request) {}
func (rh *RecipeHandler) DeleteRecipe(w http.ResponseWriter, *http.Request) {}
func (rh *RecipeHandler) RecipeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "recipeId")

		var err error
		var recipe Recipe
