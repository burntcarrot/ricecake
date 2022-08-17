<div align = "center">
    <img src="static/ricecake.png">
</div>

# ricecake

`ricecake` is a lightweight framework to build CLI applications in Go.

## Features

- Zero dependencies <sup>only uses <a href="https://github.com/spf13/pflag">pflag</a>; based on the standard library</sup>
- Support for subcommands
- POSIX-compliant flags (short/long flag styles)
- Action-based API (inspired by [urfave/cli](https://github.com/urfave/cli))
- Init Hook (hook for customizing behavior during CLI initialization)
- Auto-generated help menu!
- Custom banners
- Chainable! (see [example](./examples/chain/))
- Hide commands (see [example](./examples/hide/))
- and more!

## Installation

Installing `ricecake` is easy, just import:

```go
import "github.com/burntcarrot/ricecake"
```

## Usage

**A list of examples:**

- [Create a basic CLI with ricecake](./examples/cli/)
- [CLI with subcommands](./examples/subcommands/)
- [CLI with flags](./examples/flags/)
- [CLI built with chaining](./examples/chain/)
- [Extract extra arguments passed to CLI](./examples/extra/)
- [Hidden commands](./examples/hide/)
- [Custom banners](./examples/banner/)
- [Nested subcommands](./examples/nested/)

Here is an example on using `ricecake` to create CLIs:

```go
package main

import (
	"fmt"
	"log"

	"github.com/burntcarrot/ricecake"
)

func main() {
	// Create a new CLI.
	cli := ricecake.NewCLI("CLI", "CLI with Subcommands", "v0.1.0")

	// Set long description for the CLI.
	cli.LongDescription("This is an example CLI created with ricecake (contains subcommands).")

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

	// Run the CLI.
	err := cli.Run()
	if err != nil {
		log.Fatalf("failed to run CLI; err: %v", err)
	}
}
```

## License

`ricecake` is released under the MIT license. See [LICENSE](./LICENSE).
