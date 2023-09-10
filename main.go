package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Impossible de charger le fichier .env")
	}
	fmt.Println("Hello World")
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("No Port")
	} else {
		fmt.Println("Port: " + port)
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	errSrv := srv.ListenAndServe()
	if errSrv != nil {
		log.Fatal(errSrv)
	}
}
