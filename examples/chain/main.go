package main

import (
	"fmt"
	"log"

	"github.com/burntcarrot/ricecake"
)

func main() {
	var name string

	// Create a new CLI through chaining.
	err := ricecake.NewCLI("CLI", "CLI (chained)", "v0.1.0").
		LongDescription("This is an example CLI created through chaining.").
		StringFlagP("filename", "f", "", &name).
		Action(func() error {
			fmt.Println("I am the root command!")
			return nil
		}).Run()

	if err != nil {
		log.Fatalf("failed to run CLI; err: %v", err)
	}
}
