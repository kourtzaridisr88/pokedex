package commands

import (
	"fmt"
	"os"
)

func ExitCommand() CliCommand {
	return CliCommand{
		Name:              "exit",
		Description:       "Exit from pokedexcli",
		NumberOfArguments: 0,
		Callback: func(args []string) error {
			fmt.Println("Thank you for using pokedexcli! GoodBye!")
			os.Exit(0)
			return nil
		},
	}
}
