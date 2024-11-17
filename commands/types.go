package commands

type CliCommand struct {
	Name              string
	Description       string
	NumberOfArguments int
	Callback          func(args []string) error
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous any        `json:"previous"`
	Results  []Location `json:"results"`
}

type Pokemon struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
	BaseExp int    `json:"base_experience"`
}

type LocationArea struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokedex struct {
	Pokemons map[string]Pokemon
}

func (p *Pokedex) Catch(pokemon Pokemon) {
	p.Pokemons[pokemon.Name] = pokemon
}
