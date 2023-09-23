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
	// feed, err := urlToFeed("https://www.lemonde.fr/rss/une.xml")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(feed)

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
	// users handlers
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	// feeds handlers
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))

	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))
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
