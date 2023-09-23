package main

import (
	"fmt"
	"net/http"

	"github.com/MathieuSchaff/articlesagg/internal/auth"
	"github.com/MathieuSchaff/articlesagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Error parsing API key %v", err))
			return
		}
		user, err := cfg.DB.GetUserByAPIKEY(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Error getting user %v", err))
			return
		}
		handler(w, r, user)
	}
}
