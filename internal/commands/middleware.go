package commands

import (
	"context"
	"gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd command, user database.User) error) func(*State, command) error {

	f := func(s *State, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
	return f
}
