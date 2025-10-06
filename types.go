package main

import (
	"github.com/KubaRocks/blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}
