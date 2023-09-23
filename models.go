package main

import (
	"time"

	"github.com/MathieuSchaff/articlesagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func datababaseFeeds(databaseFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(databaseFeeds))
	for _, dbFeed := range databaseFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollow []database.FeedFollow) []FeedFollow {
	feedsFollows := make([]FeedFollow, len(dbFeedFollow))
	for _, dbFeedFollow := range dbFeedFollow {
		feedsFollows = append(feedsFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedsFollows
}

type DeleteFeedFollowParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func databaseDeleteFeedFollowParams(dbDeleteFeedFollowParams database.DeleteFeedFollowParams) DeleteFeedFollowParams {
	return DeleteFeedFollowParams{
		ID:     dbDeleteFeedFollowParams.ID,
		UserID: dbDeleteFeedFollowParams.UserID,
	}
}
