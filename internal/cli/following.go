package cli

import (
	"context"
	"fmt"

	"github.com/shubh1855/Gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {

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
