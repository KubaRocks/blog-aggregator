package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}

	for _, usr := range users {
		if usr.Name == s.cfg.CurrentUsername {
			fmt.Printf("* %v (current)\n", usr.Name)
			continue
		}
		fmt.Printf("* %v\n", usr.Name)
	}

	return nil
}
