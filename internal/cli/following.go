package cli

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	user, err := getCurrentUser(s)
	if err != nil {
		return err
	}

	follows, err := s.DB.GetFeedFollowsForUser(
		context.Background(),
		user.ID,
	)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}

	return nil
}
