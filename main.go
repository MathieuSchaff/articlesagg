package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MathieuSchaff/articlesagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Impossible de charger le fichier .env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("No Port")
	} else {
		fmt.Println("Port: " + port)
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("No DB_URL")
	}
	myDB, errDB := sql.Open("postgres", dbURL)
	if errDB != nil {
		log.Fatal(errDB)
	}
	defer myDB.Close()
	// queries will be an address that points to the location in memory where the Queries struct is stored

	queries := database.New(myDB)

	apiCfg := &apiConfig{
		DB: queries,
	}
	// creating a new router

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	r.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	errSrv := srv.ListenAndServe()
	if errSrv != nil {
		log.Fatal(errSrv)
	}
}
