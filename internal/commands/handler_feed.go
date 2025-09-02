package commands

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"gator/internal/rss"
	"time"

	"github.com/google/uuid"
)

func HandlerFetchFeed(s *State, cmd command) error {
	if len(cmd.args) != 3 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[2])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", cmd.args[2])

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func HandlerAddFeed(s *State, cmd command, user database.User) error {
	if len(cmd.args) < 4 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.name)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[2],
		Url:       cmd.args[3],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	return nil
}

func HandlerGetFeedsInfo(s *State, cmd command) error {
	if len(cmd.args) > 2 {
		return fmt.Errorf("usage: %v", cmd.name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		user, err := s.db.GetUser(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		printFeedUser(feed, user)
	}
	return nil
}

func scrapeFeeds(s *State) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	s.db.MarkFeedFetched(context.Background(), feed.ID)
	fetched_feed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("Fetching: %v\n", fetched_feed.Channel.Title)
	for _, item := range fetched_feed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})

		if err != nil {
			if err.Error() == `pq: duplicate key value violates unique constraint "posts_url_key"` {
				continue
			}
			return fmt.Errorf("failed to create post: %v", err)
		}
	}

	fmt.Printf("Fetching complete: %v\n", fetched_feed.Channel.Title)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %v\n", feed.ID)
	fmt.Printf("* CreatedAt:     %v\n", feed.CreatedAt)
	fmt.Printf("* UpdatedAt: 	 %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:      	 %v\n", feed.Name)
	fmt.Printf("* Url: 		 	 %v\n", feed.Url)
	fmt.Printf("* UserID:    	 %v\n", feed.UserID)
	fmt.Printf("* LastFetchedAt: %v\n", feed.LastFetchedAt.Time)
}

func printFeedUser(feed database.Feed, user database.User) {
	printFeed(feed)
	fmt.Printf("* User Name: 	 %v\n", user.Name)
}
