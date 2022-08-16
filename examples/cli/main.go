package main

import (
	"fmt"

	"github.com/burntcarrot/ricecake"
)

func main() {
	// Create a new CLI.
	cli := ricecake.NewCLI("CLI", "Example CLI", "v0.1.0")

	// Set long description for the CLI.
	cli.LongDescription("This is an example CLI created with ricecake.")

	// Set the action for the CLI.
	cli.Action(func() error {
		fmt.Println("Hello, ricecake!")
		return nil
	})

	// Run the CLI.
	cli.Run()
}
