package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shubh1855/Gator/internal/database"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("usage: follow <url>")
	}

	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(
		context.Background(),
		url,
	)
	if err != nil {
		return err
	}

	follow, err := s.DB.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf(
		"%s is now following %s\n",
		follow.UserName,
		follow.FeedName,
	)

	return nil
}
