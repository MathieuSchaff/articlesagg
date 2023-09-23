package main

import (
	"fmt"
	"log"
	"time"

	"github.com/MathieuSchaff/articlesagg/internal/database"
)

func startScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Starting scrapping with %v goroutines %s", concurrency, timeBetweenRequest)
	feed, err := urlToFeed("https://www.lemonde.fr/rss/une.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)
}
