package main

import "fmt"

func (c *commands) run(s *state, cmd command) error {
	h, exists := c.handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("command '%s' not registered", cmd.Name)
	}

	return h(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
