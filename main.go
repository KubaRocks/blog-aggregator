package main

import (
	"log"
	"os"

	"github.com/KubaRocks/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		cfg: &cfg,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	err = commands.run(programState, cmd)
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}
}
