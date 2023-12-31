package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MathieuSchaff/articlesagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	fmt.Printf("params r.body feed_id inhandlerCreateFeedFollow %s", params)
	defer r.Body.Close()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed follow %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed follow %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdString := chi.URLParam(r, "feedFollowID")
	feedFollowId, err := uuid.Parse(feedFollowIdString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing feed follow ID %v", err))
		return
	}
	deleteFeedFollowsParams := database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), deleteFeedFollowsParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting feed follow %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, "deleted feed follow")
}
