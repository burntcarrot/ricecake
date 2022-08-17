package main

import (
	"fmt"
	"log"

	"github.com/burntcarrot/ricecake"
)

func main() {
	// Create a new CLI.
	cli := ricecake.NewCLI("CLI", "CLI with flags", "v0.1.0")

	// Set long description for the CLI.
	cli.LongDescription("This is an example CLI created with ricecake (contains flags).")

	// -f, --file flag.
	var filename string
	cli.StringFlagP("filename", "f", "Filename", &filename)

	// Set the action for the CLI.
	cli.Action(func() error {
		fmt.Println("I am the root command!")
		fmt.Printf("-f, --file flag value: %s\n", filename)
		return nil
	})

	// Run the CLI.
	err := cli.Run()
	if err != nil {
		log.Fatalf("failed to run CLI; err: %v", err)
	}
}
