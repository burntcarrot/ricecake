package main

import (
	"fmt"
	"log"

	"github.com/burntcarrot/ricecake"
)

func main() {
	// Create a new CLI.
	cli := ricecake.NewCLI("CLI", "CLI with hidden subcommand", "v0.1.0")

	// Set long description for the CLI.
	cli.LongDescription("This is an example CLI where a subcommand is hidden.")

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

	// Hide the "greet" subcommand.
	greet.Hide()

	// Run the CLI.
	err := cli.Run()
	if err != nil {
		log.Fatalf("failed to run CLI; err: %v", err)
	}
}
