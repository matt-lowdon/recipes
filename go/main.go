package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lib/pq"

	"example.com/recipe-storage/internal/config"
)

type recipe struct {
	Name         string
	Ingredients  []string
	Instructions []string
	CookingTime  int64
	PeopleCount  int64
}

func main() {
	cfg := config.LoadConfig()

	dbConfig = pq.Config{
		Host : cfg.DB.Host,
		Port : cfg.DB.Port,
		User: cfg.DB.Port,
		Password: cfg.DB.Password,

	// pqConf = pq.Config(
	// 	url
	// )

	db, err := sql.Open("postgres", "postgres://postgres_db:5434/recipes")

	if err != nil {
		// Return is only for local development
		slog.Error("Failed to open connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Testing db connection

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":3000", router)
}

func getRecipes(w http.ResponseWriter, r *http.Request, db *sql.DB) {

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
