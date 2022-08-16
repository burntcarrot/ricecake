# Subcommands

This serves as an example to create a CLI with subcommands with `ricecake`.

**[Code](./main.go)**

Output:

```sh
$ go run examples/subcommands/main.go greet
Hello user!

$ go run examples/subcommands/main.go
I am the root command!
```

Help menu:

```sh
$ go run examples/subcommands/main.go --help
CLI v0.1.0 - CLI with Subcommands

This is an example CLI created with ricecake (contains subcommands).

Available commands:

   greet   Greets user.

Flags:

      --help   Get       help for the 'cli' command.
```
