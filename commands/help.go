package commands

import "fmt"

func HelpCommand() CliCommand {
	return CliCommand{
		Name:              "help",
		Description:       "Help!",
		NumberOfArguments: 0,
		Callback: func(args []string) error {
			text := `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Fetch Location if used multiple times would fetch the next page
explore: Explore location area and get pokemon requires argument {location-area}
`
			fmt.Println(text)
			return nil
		},
	}
}
