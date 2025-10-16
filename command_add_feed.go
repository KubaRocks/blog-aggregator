package main

import (
	"context"
	"fmt"

	"github.com/KubaRocks/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)

	if err != nil {
		return fmt.Errorf("error fetching user from database: %v", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   name,
		Url:    url,
		UserID: usr.ID,
	})

	if err != nil {
		return fmt.Errorf("error while creating an user: %v", err)
	}

	fmt.Println("New Feed added:", feed.ID, feed.Name, feed.Url)

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: usr.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error while creating a follow feed: %v", err)
	}

	fmt.Printf("Feed %s followed by %s", follow.FeedName, follow.UserName)

	return nil
}
