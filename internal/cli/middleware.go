package cli

import (
	"context"

	"github.com/shubh1855/Gator/internal/database"
)

func MiddlewareLoggedIn(
	handler func(*State, Command, database.User) error,
) func(*State, Command) error {

	return func(s *State, cmd Command) error {
		user, err := s.DB.GetUser(
			context.Background(),
			s.Config.CurrentUserName,
		)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
