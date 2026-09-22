package main

type Ingredient struct{
	Name string `json="name"`
	Quantity float64 `json="quantity"`
	Unit string `json="unit"`
}

type Recipe struct {
	Name string `json="name"`
	Ingredients []Ingredient `json="ingredients"`
	Steps []string `json="steps"`
	CookingTimeMinutes int64 `json="cooking_time_minutes"`
	Servings int64 `json="servings"`
}
