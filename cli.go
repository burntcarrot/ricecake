package ricecake

import (
	"fmt"
	"log"
	"os"
)

// CLI represents the ricecake CLI.
type CLI struct {
	version string

	rootCmd    *Command
	defaultCmd *Command

	initCommand  func(*CLI) error
	banner       func(*CLI) string
	errorHandler func(string, error) error
}

/////////////
// "Getters"
////////////

// Version returns the CLI's version.
func (c *CLI) Version() string {
	return c.version
}

// Name returns the CLI's name.
func (c *CLI) Name() string {
	return c.rootCmd.name
}

// ShortDescription returns the CLI's short description.
func (c *CLI) ShortDescription() string {
	return c.rootCmd.shortDescription
}

/////////////
// "Setters"
////////////

// SetInit calls the given function before running the CLI.
func (c *CLI) SetInit(fn func(*CLI) error) {
	c.initCommand = fn
}

// Action sets the action for the CLI.
func (c *CLI) Action(action Action) *CLI {
	c.rootCmd.Action(action)
	return c
}

// LongDescription sets the CLI's long description.
func (c *CLI) LongDescription(longdescription string) *CLI {
	c.rootCmd.LongDescription(longdescription)
	return c
}

// SetBanner sets the CLI's banner function.
func (c *CLI) SetBanner(fn func(*CLI) string) {
	c.banner = fn
}

// SetErrorHandler sets the CLI's custom error handler.
func (c *CLI) SetErrorHandler(fn func(string, error) error) {
	c.errorHandler = fn
}

////////////////////////
// Printing utilities
////////////////////////

// PrintBanner prints the CLI banner.
func (c *CLI) PrintBanner() {
	fmt.Println(c.banner(c))
	fmt.Println("")
}

// PrintHelp prints the help menu for the CLI.
func (c *CLI) PrintHelp() {
	c.rootCmd.PrintHelp()
}

///////////////////////
// Control utilities
//////////////////////

// Run runs the CLI with the given arguments.
func (c *CLI) Run(args ...string) error {
	if c.initCommand != nil {
		err := c.initCommand(c)
		if err != nil {
			return err
		}
	}

	if len(args) == 0 {
		args = os.Args[1:]
	}

	return c.rootCmd.run(args)
}

// Abort returns the given error and terminates the CLI.
func (c *CLI) Abort(err error) {
	log.Fatal(err)
	os.Exit(1)
}

//////////////
// Commands
/////////////

// AddCommand adds a command to the CLI.
func (c *CLI) AddCommand(command *Command) {
	c.rootCmd.AddCommand(command)
}

// DefaultCommand sets the given command as the default one.
func (c *CLI) DefaultCommand(defaultCmd *Command) *CLI {
	c.defaultCmd = defaultCmd
	return c
}

// NewSubCommand creates a new subcommand for the CLI.
func (c *CLI) NewSubCommand(name, description string) *Command {
	return c.rootCmd.NewSubCommand(name, description)
}

//////////////
// CLI Flags
//////////////

// BoolFlag adds a boolean flag to the root command.
func (c *CLI) BoolFlag(name, description string, variable *bool) *CLI {
	c.rootCmd.BoolFlag(name, description, variable)
	return c
}

// BoolFlagP adds a boolean flag (with shorthand) to the root command.
func (c *CLI) BoolFlagP(name, shorthand, description string, variable *bool) *CLI {
	c.rootCmd.BoolFlagP(name, shorthand, description, variable)
	return c
}

// StringFlag adds a string flag to the root command.
func (c *CLI) StringFlag(name, description string, variable *string) *CLI {
	c.rootCmd.StringFlag(name, description, variable)
	return c
}

// StringFlag adds a string flag (with shorthand) to the root command.
func (c *CLI) StringFlagP(name, shorthand, description string, variable *string) *CLI {
	c.rootCmd.StringFlagP(name, shorthand, description, variable)
	return c
}

// IntFlag adds an int flag to the root command.
func (c *CLI) IntFlag(name, description string, variable *int) *CLI {
	c.rootCmd.IntFlag(name, description, variable)
	return c
}

// IntFlagP adds an int flag (with shorthand) to the root command.
func (c *CLI) IntFlagP(name, shorthand, description string, variable *int) *CLI {
	c.rootCmd.IntFlagP(name, shorthand, description, variable)
	return c
}
