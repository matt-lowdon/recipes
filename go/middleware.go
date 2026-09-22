package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func RecipeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var recipe *Recipe
		var err error

		if articleId := chi.URLParam(r, "recipeId"); recipeId != "" {
			article, err = dbGetRecipe(recipeId)
