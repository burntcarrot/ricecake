# Hide commands

This serves as an example to hide commands in a CLI created with ricecake.

**[Code](./main.go)**

Output:

```sh
$ go run examples/hide/main.go
I am the root command!
```

Help menu:

```sh
$ go run examples/hide/main.go --help
CLI v0.1.0 - CLI with hidden subcommand

This is an example CLI where a subcommand is hidden.

Available commands:


Flags:

      --help   Get       help for the 'cli' command.
```
