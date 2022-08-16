package main

import (
	"fmt"

	"github.com/burntcarrot/ricecake"
)

func main() {
	// Create a new CLI.
	cli := ricecake.NewCLI("CLI", "CLI (extra args)", "v0.1.0")

	// Set long description for the CLI.
	cli.LongDescription("This is an example CLI which displays the extra args passed.")

	// -f, --file flag.
	var filename string
	cli.StringFlagP("filename", "f", "Filename", &filename)

	// Set the action for the CLI.
	cli.Action(func() error {
		fmt.Println("I am the root command!")
		fmt.Printf("extra args: %+v\n", cli.ExtraArgs())
		return nil
	})

	// Run the CLI.
	cli.Run()
}
