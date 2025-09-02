package commands

import (
	"context"
	"fmt"
)

func HandlerDeleteUsersTable(s *State, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}

func HandlerDeleteFeedsTable(s *State, cmd command) error {
	err := s.db.DeleteFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete feeds: %w", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}

func HandlerDeleteTables(s *State, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	err = s.db.DeleteFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete feeds: %w", err)
	}

	err = s.db.DeleteFeedFollows(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete feeds: %w", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}
