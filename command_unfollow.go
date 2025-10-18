package main

import (
	"context"
	"fmt"

	"github.com/KubaRocks/blog-aggregator/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), url)

	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}

	err = s.db.Unfollow(context.Background(), database.UnfollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow due to error: %v", err)
	}

	return nil
}
