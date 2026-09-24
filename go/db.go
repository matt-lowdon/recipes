package main

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
