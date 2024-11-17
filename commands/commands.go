package commands

func RegisterCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"explore": ExploreCommand(),
		"help":    HelpCommand(),
		"exit":    ExitCommand(),
		"map":     MapCommand(),
		"catch":   CatchCommand(),
		"pokedex": PokedexCommand(),
	}
}
