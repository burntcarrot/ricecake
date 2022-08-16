# Extra args

This serves as an example to extract extra arguments passed to a CLI created with ricecake.

**[Code](./main.go)**

Output:

```sh
$ go run examples/extra/main.go
I am the root command!
extra args: []

$ go run examples/extra/main.go a b c d
I am the root command!
extra args: [a b c d]
```

Help menu:

```sh
$ go run examples/extra/main.go --help
CLI v0.1.0 - CLI (extra args)

This is an example CLI which displays the extra args passed.

Flags:

  -f, --filename string   Filename
      --help              Get    help for the 'cli' command.
```
