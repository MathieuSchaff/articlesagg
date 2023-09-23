package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MathieuSchaff/articlesagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
		// UserID uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	defer r.Body.Close()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

// "https://www.toptal.com/developers/blog/web-front-end?utm_source=feedspot"

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting feeds %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, datababaseFeeds(feeds))
}
