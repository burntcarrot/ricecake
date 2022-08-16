# Chain

This serves as an example to create a CLI through chaining.

**[Code](./main.go)**

Output:

```sh
$ go run examples/chain/main.go
I am the root command!
```

Help menu:

```sh
$ go run examples/chain/main.go --help
CLI v0.1.0 - CLI (chained)

This is an example CLI created through chaining.

Flags:

  -f, --filename string
      --help              Get    help for the 'cli' command.
```
