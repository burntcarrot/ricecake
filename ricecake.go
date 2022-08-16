package ricecake

// NewCLI creates a new CLI.
func NewCLI(name, description, version string) *CLI {
	// Create CLI.
	cli := &CLI{
		version: version,
		banner:  defaultBanner,
	}

	// Setup root command.
	cli.rootCmd = NewCommand(name, description)
	cli.rootCmd.setCLI(cli)
	cli.rootCmd.setCommandPath("")

	return cli
}
