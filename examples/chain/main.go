package main

import (
	"fmt"

	"github.com/burntcarrot/ricecake"
)

func main() {
	var name string

	// Create a new CLI through chaining.
	ricecake.NewCLI("CLI", "CLI (chained)", "v0.1.0").
		LongDescription("This is an example CLI created through chaining.").
		StringFlagP("filename", "f", "", &name).
		Action(func() error {
			fmt.Println("I am the root command!")
			return nil
		}).Run()
}
