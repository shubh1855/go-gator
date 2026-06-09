package rss

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
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

		_, err := db.CreatePost(
			ctx,
			database.CreatePostParams{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Title:     item.Title,
				Url:       item.Link,
				Description: sql.NullString{
					String: item.Description,
					Valid:  item.Description != "",
				},
				PublishedAt: sql.NullTime{
					Time:  parsePublishedTime(item.PubDate),
					Valid: item.PubDate != "",
				},
				FeedID: feed.ID,
			},
		)

		if err != nil {

			if strings.Contains(
				err.Error(),
				"duplicate key",
			) {
				continue
			}

			fmt.Println(err)
		}
	}

	return nil
}

func parsePublishedTime(value string) time.Time {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t
		}
	}

	return time.Time{}
}
