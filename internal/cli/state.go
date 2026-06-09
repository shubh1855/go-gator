package cli

import (
	"github.com/shubh1855/go-gator/internal/config"
	"github.com/shubh1855/go-gator/internal/database"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
}

func NewState(cfg *config.Config, db *database.Queries) *State {
	return &State{
		Config: cfg,
		DB:     db,
	}
}
