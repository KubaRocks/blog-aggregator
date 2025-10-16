package main

import (
	"context"
	"fmt"

	"github.com/KubaRocks/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("failed to fetch following feeds: %v", err)
	}

	for _, follow := range following {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
