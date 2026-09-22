package main

import "net/http"

type RecipeHandler struct{
}

func (rh RecipeHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {}
func (rh RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {}
func (rh RecipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {}
func (rh RecipeHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}
func (rh RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}
