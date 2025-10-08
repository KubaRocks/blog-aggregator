package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAll(context.Background())
	if err != nil {
		return fmt.Errorf("deleting all records failed: %v", err)
	}

	return nil
}
