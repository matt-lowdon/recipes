package main

import (
	"context"
	"log/slog"
	"os"
	"errors"
	
	"github.com/jackc/pgx/v5"
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

func (pg *postgres) GetRecipe(recipeId string, ctx context.Context) (*Recipe, error) {
	var recipe *Recipe

	args := pgx.NamedArgs{
		"recipe_id" : recipeId,
	}
	
	ingredientQuery := `WITH selected_ingredients AS (
		SELECT ingredient_id, quantity, unit, preparation
		FROM recipes_api.recipe_ingredients
		WHERE recipe_id = @recipe_id),
		joined_ingredients AS (
			SELECT * FROM selected_ingredients
			JOIN recipes_api.ingredients
			ON selected_ingredients.ingredient_id = ingredients.id
		)
		SELECT name, quantity, unit, preparartion FROM joined_ingredients;`


	instructionQuery := `SELECT step_text FROM recipes_api.recipe_instructions
	WHERE recipe_id = @recipe_id ORDER BY position;`
	
	recipeQuery := `SELECT name, servings, cook_time_minutes FROM recipes_api.recipes
	WHERE id = @recipe_id;`

	rows, err := pg.db.Exec(ctx, ingredientQuery, args)
	if err != nil {
		return nil, errors.New("issue retrieving ingredients", "error", err)
	}

	ingredients = pgx.CollectRows(rows, pgx.RowToStructByName[Ingredient])

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
