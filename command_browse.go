package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/KubaRocks/blog-aggregator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 && cmd.Args[0] != "" {
		if n, err := strconv.Atoi(cmd.Args[0]); err == nil && n > 0 {
			limit = n
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to get posts from user: %v", err)
	}

	for _, post := range posts {
		fmt.Printf("%s\n", post.Title.String)
		fmt.Printf("Published At: %v\n", post.PublishedAt)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Println("-=-=-=-=-=-")
		fmt.Println(post.Description.String)
		fmt.Println("----------------------------------------------------------------------")
	}

	return nil
}
