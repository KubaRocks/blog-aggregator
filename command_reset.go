package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	// feeds will be deleted by cascade
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("deleting all users records failed: %v", err)
	}

	return nil
}
