package ricecake

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

// Action represents a function that gets called when a command is used.
type Action func() error

// Command represents a CLI command.
type Command struct {
	name             string
	commandPath      string
	shortDescription string
	longDescription  string

	subCommands    []*Command
	subCommandsMap map[string]*Command

	cli    *CLI
	flags  *flag.FlagSet
	action Action

	flagCount               int
	longestSubcommandLength int
	helpFlag                bool
	hidden                  bool
}

////////////////////////////////////////
// Command registration and creation
////////////////////////////////////////

// NewCommand creates a new command.
func NewCommand(name, description string) *Command {
	command := &Command{
		name:             name,
		shortDescription: description,
		subCommandsMap:   make(map[string]*Command),
		hidden:           false,
	}

	return command
}

// NewSubCommand creates a new subcommand.
// NOTE: NewSubCommand registers a subcommand to the root command on creation.
func (c *Command) NewSubCommand(name, description string) *Command {
	command := NewCommand(name, description)
	c.AddCommand(command)
	return command
}

// AddCommand adds a subcommand to the command.
func (c *Command) AddCommand(command *Command) {
	// Setup CLI and command path
	command.setCLI(c.cli)
	command.setCommandPath(c.commandPath)

	// Register subcommand.
	name := command.name
	c.subCommands = append(c.subCommands, command)
	c.subCommandsMap[name] = command

	// Update longestSubcommandLength on the go.
	if len(name) > c.longestSubcommandLength {
		c.longestSubcommandLength = len(name)
	}
}

/////////////
// "Setters"
/////////////

// setCommandPath sets the command path for the command.
func (c *Command) setCommandPath(commandPath string) {
	// Set up command path.
	if commandPath != "" {
		c.commandPath += commandPath + " "
	}

	c.commandPath += c.name

	c.flags = flag.NewFlagSet(c.commandPath, flag.ContinueOnError)
	c.BoolFlag("help", fmt.Sprintf("Get	 help for the '%s' command.", strings.ToLower(c.commandPath)), &c.helpFlag)
}

// setCLI sets the CLI for the command.
func (c *Command) setCLI(cli *CLI) {
	c.cli = cli
}

// LongDescription sets the long description for the command.
func (c *Command) LongDescription(longDescription string) *Command {
	c.longDescription = longDescription
	return c
}

/////////////////////
// Common utilities
/////////////////////

// parseFlags parses the given flags.
func (c *Command) parseFlags(args []string) error {
	tmp := os.Stderr
	os.Stderr = nil

	// Parse flag definitions from the argument list.
	err := c.flags.Parse(args)

	os.Stderr = tmp
	return err
}

// Run executes the command with the given arguments.
func (c *Command) run(args []string) error {
	if len(args) > 0 {
		subcommand := c.subCommandsMap[args[0]]
		if subcommand != nil {
			return subcommand.run(args[1:])
		}

		err := c.parseFlags(args)
		if err != nil {
			if c.cli.errorHandler != nil {
				// Use the error handler to handle errors.
				return c.cli.errorHandler(c.commandPath, err)
			}

			return fmt.Errorf("error: %s\nSee '%s --help' for usage", err, c.commandPath)
		}

		if c.helpFlag {
			c.PrintHelp()
			return nil
		}
	}

	// If there is an action associated with this command, run the action.
	if c.action != nil {
		return c.action()
	}

	// Check the default command.
	// This is useful for scenarios where there are no subcommands defined.
	if c.cli.defaultCmd != nil {
		// This is an essential check.
		// If the default command is similar to the current command, executing it causes recursion.
		if c.cli.defaultCmd != c {
			// only run default command if no args passed
			if len(args) == 0 {
				return c.cli.defaultCmd.run(args)
			}
		}
	}

	// If nothing works, just print the help menu.
	c.PrintHelp()

	return nil
}

// PrintHelp generates the help text for the command.
func (c *Command) PrintHelp() {
	// Use banner.
	c.cli.PrintBanner()

	// Set command title.
	commandTitle := c.commandPath

	// Add short description to command description.
	if c.shortDescription != "" {
		commandTitle += " - " + c.shortDescription
	}

	// Ignore root command.
	if c.commandPath != c.name {
		fmt.Println(commandTitle)
	}

	// Add long description to command description.
	if c.longDescription != "" {
		fmt.Println(c.longDescription + "\n")
	}

	// Generate subcommands list.
	if len(c.subCommands) > 0 {
		fmt.Println("Available commands:")
		fmt.Println("")

		for _, subcommand := range c.subCommands {
			if subcommand.isHidden() {
				continue
			}

			// Generate spacing.
			spacing := strings.Repeat(" ", 3+c.longestSubcommandLength-len(subcommand.name))
			isDefault := ""

			// Indicate that the subcommand is the default command.
			if subcommand.isDefaultCommand() {
				isDefault = "[default]"
			}

			// Print the subcommand information.
			fmt.Printf("   %s%s%s %s\n", subcommand.name, spacing, subcommand.shortDescription, isDefault)
		}

		fmt.Println("")
	}

	// Generate flag list.
	if c.flagCount > 0 {
		// Set header.
		fmt.Println("Flags:")
		fmt.Println("")

		c.flags.SetOutput(os.Stdout)
		c.flags.PrintDefaults()
		c.flags.SetOutput(os.Stderr)

	}

	// Terminal etiquette. End with a newline.
	fmt.Println("")
}

////////////////////////
// Command behaviors
///////////////////////

// isDefaultCommand checks if the command is the default command for the CLI.
func (c *Command) isDefaultCommand() bool {
	return c.cli.defaultCmd == c
}

// isHidden checks if the command is a hidden command.
func (c *Command) isHidden() bool {
	return c.hidden
}

// Hide hides the command from the help menu.
func (c *Command) Hide() {
	c.hidden = true
}

/////////////
// Action
////////////

// Action sets an action for the command.
func (c *Command) Action(action Action) *Command {
	c.action = action
	return c
}

/////////////
// Others
////////////

// ExtraArgs returns the non-flag arguments passed to the command.
func (c *Command) ExtraArgs() []string {
	return c.flags.Args()
}

////////////
// Flags
///////////

// BoolFlag adds a boolean flag to the command.
func (c *Command) BoolFlag(name, description string, variable *bool) *Command {
	c.flags.BoolVar(variable, name, *variable, description)
	c.flagCount++
	return c
}

// BoolFlagP adds a boolean flag (with shorthand) to the command.
func (c *Command) BoolFlagP(name, shorthand, description string, variable *bool) *Command {
	c.flags.BoolVarP(variable, name, shorthand, *variable, description)
	c.flagCount++
	return c
}

// StringFlag adds a string flag to the command.
func (c *Command) StringFlag(name, description string, variable *string) *Command {
	c.flags.StringVar(variable, name, *variable, description)
	c.flagCount++
	return c
}

// StringFlag adds a string flag (with shorthand) to the command.
func (c *Command) StringFlagP(name, shorthand, description string, variable *string) *Command {
	c.flags.StringVarP(variable, name, shorthand, *variable, description)
	c.flagCount++
	return c
}

// IntFlag adds an int flag to the command.
func (c *Command) IntFlag(name, description string, variable *int) *Command {
	c.flags.IntVar(variable, name, *variable, description)
	c.flagCount++
	return c
}

// IntFlag adds an int flag (with shorthand) to the command.
func (c *Command) IntFlagP(name, shorthand, description string, variable *int) *Command {
	c.flags.IntVarP(variable, name, shorthand, *variable, description)
	c.flagCount++
	return c
}
