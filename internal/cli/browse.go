package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shubh1855/go-gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := int32(2)

	if len(cmd.Args) == 2 {
		val, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return err
		}

		if val < 0 {
			return fmt.Errorf("limit  must be non-negative")
		}

		limit = int32(val)
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s\n%s\n", post.Title, post.Url)
	}

	return nil
}
