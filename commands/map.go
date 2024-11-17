package commands

import (
	"encoding/json"
	"fmt"

	"github.com/kourtzaridisr88/pokedexcli/client"
	"github.com/kourtzaridisr88/pokedexcli/pokecache"
)

func callEndpointAndCache() ([]Location, error) {
	cache := pokecache.CacheInstance()

	if cachedData, found := cache.Get("location"); found {
		var data []Location
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return nil, fmt.Errorf("error unmarshaling cached data: %v", err)
		}
		fmt.Println("FETCHED FROM CACHE")
		return data, nil
	}

	res, err := client.GetEndpoint("location")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fmt.Println("FETCHED FROM API")

	var locRes LocationResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locRes); err != nil {
		return nil, err
	}

	// Store the fetched results in the cache
	data, err := json.Marshal(locRes.Results)
	if err != nil {
		return nil, fmt.Errorf("error marshaling response for cache: %v", err)
	}
	cache.Add("location", data, 10*60)

	return locRes.Results, nil
}

func MapCommand() CliCommand {
	return CliCommand{
		Name:              "map",
		Description:       "Displays the names of 20 location areas in the Pokemon world",
		NumberOfArguments: 0,
		Callback: func(args []string) error {
			results, err := callEndpointAndCache()
			if err != nil {
				return err
			}

			for _, element := range results {
				fmt.Println(element.Name)
			}

			return nil
		},
	}
}
