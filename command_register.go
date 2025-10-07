package main

import (
	"context"
	"fmt"
	"time"

	"github.com/KubaRocks/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	usr, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("user already exists: %s", name)
	}

	s.cfg.SetUser(name)
	fmt.Printf("User %s was created: %v", name, usr)

	return nil
}
