package cli

import (
	"context"

	"github.com/shubh1855/Gator/internal/database"
)

func getCurrentUser(s *State) (database.User, error) {
	return s.DB.GetUser(
		context.Background(),
		s.Config.CurrentUserName,
	)
}
