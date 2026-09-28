package main

import (
	"context"
	"log/slog"
	"os"
	
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

func newPg(ctx context.Context, connString string) (*postgres, error) {
	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	return &postgres{db}, nil
}

func (pg *postgres) GetRecipe(recipeId string) (*Recipe, error) {
	var recipe *Recipe
	


var recipesDb = []*Recipe{
	{
		ID: "1",
		Name: "Test",
		Ingredients : []Ingredient{
			{
				Name : "Onion",
				Quantity : 1,
				Unit : "whole",
			},
		},
		Instructions : []string{"Slice the onion", "Fry the onion"},
		CookTimeMinutes : 15,
		Servings : 2,
	},
}
