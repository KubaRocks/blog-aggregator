package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/KubaRocks/blog-aggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed parsing time_between_reqs argument: %v", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
		log.Printf("waiting %s for next fetch", timeBetweenRequests)
	}
}

func scrapeFeeds(s *state) error {
	dbFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed fetching next feed: %v", err)
	}
	log.Println("Found a feed to fetch!")
	err = scrapeFeed(s.db, dbFeed)

	if err != nil {
		return fmt.Errorf("error scraping feed: %v", err)
	}

	return nil
}

func scrapeFeed(db *database.Queries, dbFeed database.Feed) error {
	feed, err := fetchFeed(context.Background(), dbFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed from web: %v", err)
	}

	err = db.MarkFeedFetched(context.Background(), dbFeed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %v", err)
	}

	fmt.Printf("\n----------------\nPrinting Items from %s\n----------------\n\n", feed.Channel.Title)

	for _, feedItem := range feed.Channel.Item {
		fmt.Printf("* %s\n", feedItem.Title)
	}
	log.Printf("Feed %s collected, %v posts found", dbFeed.Name, len(feed.Channel.Item))

	return nil
}
