package main

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/lib/pq"
	// "example.com/recipe-storage/internal/config"
)

func main() {
	// cfg := config.LoadConfig()

	// dbConfig = pq.Config{
	// 	Host : cfg.DB.Host,
	// 	Port : cfg.DB.Port,
	// 	User: cfg.DB.Port,
	// 	Password: cfg.DB.Password,

	// ParseUint returns unit64, which needs converting for pq.Config
	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16)
	if err != nil {
		slog.Warn("No port env var found. Defaulting to 5434")
		port = uint64(5434)
	}

	pqConf := pq.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     uint16(port),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	db, err := connectToDb(pqConf)
	if err != nil {
		os.Exit(1)
	}

	defer db.Close()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	http.ListenAndServe(":3000", router)
}

func connectToDb(config pq.Config) (*sql.DB, error) {
	slog.Info("Connecting to database")
	c, err := pq.NewConnectorConfig(config)

	if err != nil {
		slog.Error("pq connector error", "error", err)
		return nil, err
	}

	db := sql.OpenDB(c)
	err = db.Ping()
	if err != nil {
		slog.Error("Database connection error", "error", err)
		return nil, err
	}

	return db, nil
}

func RecipeRouter() chi.Router {
	r := chi.NewRouter()
	recipeHandler := RecipeHandler{}
	r.Get("/", recipeHandler.ListRecipes)
	r.Post("/", recipeHandler.CreateRecipe)
	r.Route("/{recipeId}", func (r chi.Router) {
		r.Use(
