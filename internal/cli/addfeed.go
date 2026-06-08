package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shubh1855/Gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("usage: addfeed <name> <url>")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := getCurrentUser(s)
	if err != nil {
		return err
	}

	feed, err := s.DB.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
			Url:       url,
			UserID:    user.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
