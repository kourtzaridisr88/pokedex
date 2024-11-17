package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/kourtzaridisr88/pokedexcli/client"
)

func callPokemonDetails(pokemonName string) (Pokemon, error) {
	endpoint := "pokemon/" + pokemonName
	var pokemon Pokemon

	res, err := client.GetEndpoint(endpoint)
	if err != nil {
		return pokemon, err
	}

	defer res.Body.Close()

	if res.StatusCode >= 404 {
		return pokemon, errors.New("pokemon not found")
	}

	if res.StatusCode >= 400 {
		return pokemon, errors.New("something went wrong")
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&pokemon); err != nil {
		return pokemon, err
	}

	return pokemon, nil
}

func CatchCommand() CliCommand {
	return CliCommand{
		Name:              "catch",
		Description:       "Catch a pokemon",
		NumberOfArguments: 1,
		Callback: func(args []string) error {
			pokemonName := args[0]
			pokemon, err := callPokemonDetails(pokemonName)
			if err != nil {
				return err
			}

			fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
			fmt.Println(pokemon)

			catchChance := max(5, 100-pokemon.BaseExp/2)

			if catchChance >= rand.IntN(100) {
				fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
				fmt.Printf("%v was caught!\n", pokemonName)
				PokedexInstance().Catch(pokemon)
			} else {
				fmt.Printf("%v escaped!\n", pokemonName)
			}

			return nil
		},
	}
}
