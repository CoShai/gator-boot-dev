package commands

import (
	"errors"
	"gator/internal/config"
	"gator/internal/database"
)

type Commands struct {
	cmdList map[string]func(*State, command) error
}

type State struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}

func (c *Commands) Run(s *State, cmd command) error {
	f, ok := c.cmdList[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State, command) error) {
	c.cmdList[name] = f
}

func CreateState(cfg *config.Config, db *database.Queries) *State {
	state := State{
		config: cfg,
		db:     db,
	}
	return &state
}

func GetCommands() *Commands {
	cmds := Commands{
		cmdList: make(map[string]func(*State, command) error),
	}
	return &cmds
}

func GetCommand(name string, args []string) command {
	return command{
		name: name,
		args: args,
	}
}
