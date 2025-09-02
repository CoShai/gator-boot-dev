package commands

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func HandlerBrowse(s *State, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 3 {
		value, err := strconv.ParseInt(cmd.args[2], 10, 8)
		if err != nil {
			return fmt.Errorf("failed to parse string to int: %v", err)
		}
		limit = int(value)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return fmt.Errorf("failed to retieve posts from database: %v", err)
	}

	for i, post := range posts {
		if i == limit {
			return nil
		}

		fmt.Println("======================")
		fmt.Printf("Title        : %v\n", post.Title)
		fmt.Printf("Published At : %v\n", post.PublishedAt)
		fmt.Printf("URL          : %v\n", post.Url)
		fmt.Printf("Description  : %v\n", post.Description.String)
		fmt.Println("======================")
	}
	return nil
}
