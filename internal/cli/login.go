package cli

import (
	"context"
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("login requires exactly one username")
	}

	name := cmd.Args[0]

	_, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		return errors.New("user does not exist")
	}

	err = s.Config.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("Logged in as %s\n", name)
	return nil
}
