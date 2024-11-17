package commands

import "fmt"

var pokedexInstance *Pokedex

func PokedexInstance() *Pokedex {
	if pokedexInstance == nil {
		pokedexInstance = &Pokedex{
			Pokemons: make(map[string]Pokemon),
		}
	}

	return pokedexInstance
}

func PokedexCommand() CliCommand {
	return CliCommand{
		Name:              "pokedex",
		Description:       "view your pokedex",
		NumberOfArguments: 0,
		Callback: func(args []string) error {
			fmt.Println("Your Pokedex: ")
			for _, pokemon := range PokedexInstance().Pokemons {
				fmt.Printf("- %v\n", pokemon.Name)
			}

			return nil
		},
	}
}
