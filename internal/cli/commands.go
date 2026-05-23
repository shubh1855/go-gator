package cli

import "fmt"

type Command struct {
	Name string
	Args []string
}

type Handler func(*State, Command) error

type Commands struct {
	handlers map[string]Handler
}

func NewCommands() *Commands {
	return &Commands{
		handlers: make(map[string]Handler),
	}
}

func (c *Commands) Register(name string, h Handler) {
	c.handlers[name] = h
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return handler(s, cmd)
}
