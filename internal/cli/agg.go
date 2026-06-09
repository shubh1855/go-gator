package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/shubh1855/go-gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("usage: agg <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feed every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		err := rss.ScrapeFeeds(context.Background(), s.DB)
		if err != nil {
			fmt.Println(err)
		}
	}
}
