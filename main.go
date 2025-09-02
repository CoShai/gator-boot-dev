package main

import (
	"database/sql"
	"gator/internal/commands"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	dbQueries := database.New(db)
	state := commands.CreateState(&cfg, dbQueries)

	cmds := commands.GetCommands()
	cmds.Register("help", commands.HandlerHelp)
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerDeleteTables)
	cmds.Register("users", commands.HandlerGetUsers)
	cmds.Register("agg", commands.HandlerFetchFeed)
	cmds.Register("addfeed", commands.MiddlewareLoggedIn(commands.HandlerAddFeed))
	cmds.Register("feeds", commands.HandlerGetFeedsInfo)
	cmds.Register("follow", commands.MiddlewareLoggedIn(commands.HandlerAddFeedFollow))
	cmds.Register("following", commands.MiddlewareLoggedIn(commands.HandlerGetFeedFollowingForUser))
	cmds.Register("unfollow", commands.MiddlewareLoggedIn(commands.HandlerDeleteFeedFollow))
	cmds.Register("browse", commands.MiddlewareLoggedIn(commands.HandlerBrowse))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := commands.GetCommand(os.Args[1], os.Args)

	err = cmds.Run(state, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
