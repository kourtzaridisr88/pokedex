package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kourtzaridisr88/pokedexcli/client"
)

func callLocationAreaEndpointAndCache(area string) (LocationArea, error) {
	endpoint := "location-area/" + area
	var locationArea LocationArea

	res, err := client.GetEndpoint(endpoint)
	if err != nil {
		return locationArea, err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return locationArea, errors.New("something went wrong")
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationArea); err != nil {
		return locationArea, err
	}

	return locationArea, nil
}

func ExploreCommand() CliCommand {
	return CliCommand{
		Name:              "explore",
		Description:       "Find pokemenos in location area",
		NumberOfArguments: 1,
		Callback: func(args []string) error {
			results, err := callLocationAreaEndpointAndCache(args[0])
			if err != nil {
				return err
			}

			for _, element := range results.PokemonEncounters {
				fmt.Println(element.Pokemon.Name)
			}

			return nil
		},
	}
}
