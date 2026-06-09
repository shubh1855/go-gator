package rss

import (
	"context"
	"fmt"

	"github.com/shubh1855/Gator/internal/database"
)

func ScrapeFeeds(ctx context.Context, db *database.Queries) error {

	feed, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	err = db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Fetching %s\n", feed.Name)

	rssFeed, err := FetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
