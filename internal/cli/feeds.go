package cli

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf(
			"Name: %s\nURL: %s\nUser: %s\n\n",
			feed.Name,
			feed.Url,
			feed.UserName,
		)
	}

	return nil
}
