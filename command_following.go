package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("failed to fetch following feeds: %v", err)
	}

	for _, follow := range following {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
