package main

import (
	"github.com/KubaRocks/blog-aggregator/internal/config"
	"github.com/KubaRocks/blog-aggregator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}
