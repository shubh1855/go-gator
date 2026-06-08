package cli

import (
	"context"
	"errors"

	"github.com/shubh1855/Gator/internal/database"
)

func HandlerUnfollow(
	s *State,
	cmd Command,
	user database.User,
) error {

	if len(cmd.Args) != 1 {
		return errors.New("usage: unfollow <url>")
	}

	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(
		context.Background(),
		url,
	)
	if err != nil {
		return err
	}

	err = s.DB.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
