package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kourtzaridisr88/pokedexcli/commands"
)

func main() {
	PrintWelcomeMsg()

	commands := commands.RegisterCommands()
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)

	for {
		PrintPrompt()
		reader.Scan()
		input := reader.Text()

		inputWords := strings.Fields(input)

		commandName := inputWords[0]
		arguments := inputWords[1:]

		cmd, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command. Please try again.")
			continue
		}

		if cmd.NumberOfArguments != len(arguments) {
			fmt.Printf("Command %v requires arguments. Use help to learn more \n", cmd.Name)
			continue
		}

		// Execute the command's callback
		err := cmd.Callback(arguments)
		if err != nil {
			fmt.Printf("Error executing command '%s': %v\n", input, err)
		}
	}
}

func PrintWelcomeMsg() {
	fmt.Println("Welcome to Pokedex cli!")
}

func PrintPrompt() {
	fmt.Printf("pokedex > ")
}
