package main

import(
	"net/http"
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type RecipeHandler struct{
}

func (rh *RecipeHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, RecipeListResponse()); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (rh *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := r.Context().Value("recipe").(*Recipe)
	if err := render.Render(w, r, &RecipeResponse{Recipe : recipe}); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (rh *RecipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	data := &RecipeRequest{}
	err := render.Bind(r, data); if err != nil{
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	recipe := data.Recipe
	dbNewRecipe(recipe)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &RecipeResponse{Recipe : recipe})
}

func (rh *RecipeHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := r.Context().Value("recipe").(*Recipe)
	data := &RecipeRequest{}

	err := render.Bind(r, data); if err != nil{
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	recipe = data.Recipe
	dbUpdateRecipe(recipe.ID, recipe)
	render.Render(w, r, &RecipeResponse{Recipe: recipe})
}

func (rh *RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}
func (rh *RecipeHandler) RecipeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var recipe *Recipe

		if recipeId := chi.URLParam(r, "recipeId"); recipeId != "" {
			recipe, err = dbGetRecipe(recipeId)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}

		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "recipe", recipe)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//Code supporting RecipeHandler.ListRecipes
func RecipeListResponse() []render.Renderer {
	list := []render.Renderer{}

	for _, recipe := range recipesDb {
		list = append(list, &RecipeResponse{Recipe: recipe})
	}

	return list
}
