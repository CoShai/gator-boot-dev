package commands

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func HandlerLogin(s *State, cmd command) error {
	if len(cmd.args) < 3 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	user, err := s.db.GetUserByName(context.Background(), cmd.args[2])
	if err != nil {
		return errors.New("couldn't find user")
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")

	return nil
}

func HandlerRegister(s *State, cmd command) error {
	if len(cmd.args) < 3 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[2],
	})

	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully")
	printUser(user)
	return nil
}

func HandlerGetUsers(s *State, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return errors.New("couldn't get users")
	}

	for _, user := range users {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}

	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf("* id:		 %v\n", user.ID)
	fmt.Printf("* CreatedAt: %v\n", user.CreatedAt)
	fmt.Printf("* UpdatedAt: %v\n", user.UpdatedAt)
	fmt.Printf("* Name:	 	 %v\n", user.Name)
}
