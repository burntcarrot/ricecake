package main

import (
	"fmt"
	"log"

	"github.com/burntcarrot/ricecake"
)

func customBanner(cli *ricecake.CLI) string {
	banner := `
	███                                       █████
	░░░                                       ░░███
████████  ████   ██████   ██████   ██████   ██████   ░███ █████  ██████
░░███░░███░░███  ███░░███ ███░░███ ███░░███ ░░░░░███  ░███░░███  ███░░███
░███ ░░░  ░███ ░███ ░░░ ░███████ ░███ ░░░   ███████  ░██████░  ░███████
░███      ░███ ░███  ███░███░░░  ░███  ███ ███░░███  ░███░░███ ░███░░░
█████     █████░░██████ ░░██████ ░░██████ ░░████████ ████ █████░░██████
░░░░░     ░░░░░  ░░░░░░   ░░░░░░   ░░░░░░   ░░░░░░░░ ░░░░ ░░░░░  ░░░░░░
	`
	banner += fmt.Sprintf(" %s - %s", cli.Version(), cli.ShortDescription())
	return banner
}

func main() {
	// Create a new CLI.
	cli := ricecake.NewCLI("CLI", "CLI with custom banner", "v0.1.0")

	// Set long description for the CLI.
	cli.LongDescription("This is an example CLI created with a custom banner.")

	// Set custom banner.
	cli.SetBanner(customBanner)

	// Set the action for the CLI.
	cli.Action(func() error {
		fmt.Println("I am the root command!")
		return nil
	})

	// Run the CLI.
	err := cli.Run()
	if err != nil {
		log.Fatalf("failed to run CLI; err: %v", err)
	}
}
