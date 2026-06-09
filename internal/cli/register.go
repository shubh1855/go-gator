package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shubh1855/go-gator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("register requires exactly one username")
	}

	name := cmd.Args[0]

	user, err := s.DB.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
		},
	)
	if err != nil {
		return err
	}

	if err := s.Config.SetUser(name); err != nil {
		return err
	}

	fmt.Printf("User %s created\n", name)
	fmt.Println(user)

	return nil
}
