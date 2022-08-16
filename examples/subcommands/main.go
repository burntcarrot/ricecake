package main

import (
	"fmt"

	"github.com/burntcarrot/ricecake"
)

func main() {
	// Create a new CLI.
	cli := ricecake.NewCLI("CLI", "CLI with Subcommands", "v0.1.0")

	// Set long description for the CLI.
	cli.LongDescription("This is an example CLI created with ricecake (contains subcommands).")

	// Set the action for the CLI.
	cli.Action(func() error {
		fmt.Println("I am the root command!")
		return nil
	})

	// Create a new "greet" subcommand.
	greet := cli.NewSubCommand("greet", "Greets user.")

	// Set the action for the "greet" subcommand.
	greet.Action(func() error {
		fmt.Println("Hello user!")
		return nil
	})

	// Run the CLI.
	cli.Run()
}
