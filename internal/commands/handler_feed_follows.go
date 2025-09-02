package commands

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func HandlerAddFeedFollow(s *State, cmd command, user database.User) error {
	if len(cmd.args) != 3 {
		return fmt.Errorf("usage: %v <url>", cmd.name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[2])
	if err != nil {
		return err
	}

	s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	fmt.Println("Feed followed successfully")
	return nil
}

func HandlerGetFeedFollowingForUser(s *State, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %v", cmd.name)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("user currently dont follow feeds")
		return nil
	}

	for _, feed := range feeds {
		fmt.Printf("%v\n", feed.FeedName)
	}

	return nil
}

func HandlerDeleteFeedFollow(s *State, cmd command, user database.User) error {
	if len(cmd.args) != 3 {
		return fmt.Errorf("usage: %v <feed_url>", cmd.name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[2])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID, FeedID: feed.ID})

	if err != nil {
		return err
	}

	fmt.Printf("%s unfollowed successfully!\n", feed.Name)
	return nil
}
