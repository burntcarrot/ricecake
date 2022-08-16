# Flags

This serves as an example to create a CLI with flags.

**[Code](./main.go)**

Output:

```sh
$ go run examples/flags/main.go -f hello.go
I am the root command!
-f, --file flag value: hello.go
```

Help menu:

```sh
$ go run examples/flags/main.go --help
CLI v0.1.0 - CLI with flags

This is an example CLI created with ricecake (contains flags).

Flags:

  -f, --filename string   Filename
      --help              Get    help for the 'cli' command.
```
