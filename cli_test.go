package ricecake

import (
	"errors"
	"fmt"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("SetInit", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		cli.SetInit(func(c *CLI) error {
			return nil
		})
		cli.Run()
	})

	t.Run("SetInit with error", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		cli.SetInit(func(c *CLI) error {
			return errors.New("test error")
		})
		cli.Run()
	})

	t.Run("SetBanner", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		cli.SetBanner(func(c *CLI) string {
			return "custom banner"
		})
	})

	t.Run("handle error with SetErrorHandler", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		cli.SetErrorHandler(func(s string, err error) error {
			return fmt.Errorf("command = %s, encountered error: %v", s, err)
		})
		cli.Run("-badflag")
	})

	t.Run("handle error", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		cli.Run("-badflag")
	})

	t.Run("subcommand", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		greet := cli.NewSubCommand("greet", "Greet subcommand")
		greet.Action(func() error {
			return nil
		})

		cli.Run("test greet")
	})

	t.Run("default subcommand", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		greet := cli.NewSubCommand("greet", "Greet subcommand")
		greet.Action(func() error {
			return nil
		})
		cli.DefaultCommand(greet)

		cli.PrintHelp()
		cli.Run()
	})

	t.Run("hidden subcommand", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		greet := cli.NewSubCommand("greet", "Greet subcommand")
		greet.Action(func() error {
			return nil
		})
		greet.Hide()

		cli.PrintHelp()
	})

	t.Run("flags", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		greet := cli.NewSubCommand("greet", "Greet subcommand")
		greet.Action(func() error {
			return nil
		})

		// The flag declaration done here is not idiomatic. This is meant for testing purposes.
		// There shouldn't be multiple flags tied to a single purpose.

		var display bool
		var displayShort bool
		var filename string
		var filenameShort string
		var count int
		var countShort int

		cli.BoolFlag("display", "Display", &display)
		cli.StringFlag("filename", "Filename", &filename)
		cli.IntFlag("count", "Count", &count)

		cli.BoolFlagP("display-short", "d", "Display", &displayShort)
		cli.StringFlagP("filename-short", "f", "Filename", &filenameShort)
		cli.IntFlagP("count-short", "c", "Count", &countShort)

		cli.PrintHelp()
	})

	t.Run("add command", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		greet := &Command{
			name:             "greet",
			shortDescription: "Greet user.",
		}
		cli.AddCommand(greet)
	})

	t.Run("ExtraArgs", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		cli.Action(func() error {
			fmt.Printf("extra args: %+v\n", cli.ExtraArgs())
			return nil
		})
		cli.Run()
	})
}
