# Nested subcommands

This serves as an example to create a CLI with nested subcommands with `ricecake`.

**[Code](./main.go)**

Output:

```sh
$ go run examples/nested/main.go greet morning
Good morning, user!
```

Help menu:

```sh
$ go run examples/nested/main.go --help
CLI v0.1.0 - CLI with nested subcommands

This is an example CLI created with ricecake (contains nested subcommands).

Available commands:

   greet   Greets user.

Flags:

      --help   Get       help for the 'cli' command.
```

```
$ go run examples/nested/main.go greet --help
CLI v0.1.0 - CLI with nested subcommands

CLI greet - Greets user.
Available commands:

   morning   Greets user with a good morning message.

Flags:

      --help   Get       help for the 'cli greet' command.
```
