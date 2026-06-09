package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shubh1855/Gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := int32(2)

	if len(cmd.Args) == 2 {
		val, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
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
