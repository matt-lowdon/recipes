package main

import(
	"net/http"
	"errors"

	"github.com/go-chi/render"
)

type Ingredient struct{
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	Unit string `json:"unit"`
}

type Recipe struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Instructions []string `json:"instructions"`
	CookTimeMinutes int `json:"cook_time_minutes"`
	Servings int `json:"servings"`
}

//render.Binder implementation for recipes model
type RecipeRequest struct{
	*Recipe
}

func (rr *RecipeRequest) Bind(r *http.Request) error {
	if rr.Recipe == nil {
		return errors.New("missing required Recipe fields")
	}

	return nil
}

//render.Renderer implementation for recipes model
type RecipeResponse struct{
	*Recipe
}

func (r *RecipeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

//Struct and method for render.Render interface. Plus calling functions.
type ErrResponse struct {
	Err error `json:"-"`
	HTTPStatusCode int `json:"-"`

	StatusText string `json:"status"`
	AppCode int64 `json:"code,omitempty"`
	ErrorText string `json:"error,omitempty"`
}

func (e *ErrResponse)  Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err: err,
		HTTPStatusCode : 400,
		StatusText : "Invalid request.",
		ErrorText : err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err : err,
		HTTPStatusCode : 422,
		StatusText : "Error rendering response.",
		ErrorText : err.Error()
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode : 404, StatusText : "Resource not found."}
