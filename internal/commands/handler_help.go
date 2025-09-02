package commands

import "fmt"

func HandlerHelp(s *State, cmd command) error {
	help := make(map[string]string)

	help["help"] = "show command list"
	help["login"] = "login user"
	help["register"] = "register user"
	help["reset"] = "reset database"
	help["users"] = "show users list"
	help["agg"] = "aggerator"
	help["addfeed"] = "add feed"
	help["feeds"] = "show feeds list"
	help["follow"] = "follow feed"
	help["following"] = "show feeds current user follow"
	help["unfollow"] = "unfollow feed"
	help["browse"] = "browse feed"

	fmt.Println()
	for key, value := range help {
		fmt.Printf("%v : %v\n", key, value)
	}

	return nil
}
