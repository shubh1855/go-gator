package cli

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.DB.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Database reset successfully")
	return nil
}
