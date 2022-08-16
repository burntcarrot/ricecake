# Banner

This serves as an example to create a CLI with a custom banner.

**[Code](./main.go)**

Output:

```sh
$ go run examples/banner/main.go
I am the root command!
```

Help menu:

```sh
$ go run examples/banner/main.go --help

        ███                                       █████
        ░░░                                       ░░███
████████  ████   ██████   ██████   ██████   ██████   ░███ █████  ██████
░░███░░███░░███  ███░░███ ███░░███ ███░░███ ░░░░░███  ░███░░███  ███░░███
░███ ░░░  ░███ ░███ ░░░ ░███████ ░███ ░░░   ███████  ░██████░  ░███████
░███      ░███ ░███  ███░███░░░  ░███  ███ ███░░███  ░███░░███ ░███░░░
█████     █████░░██████ ░░██████ ░░██████ ░░████████ ████ █████░░██████
░░░░░     ░░░░░  ░░░░░░   ░░░░░░   ░░░░░░   ░░░░░░░░ ░░░░ ░░░░░  ░░░░░░
         v0.1.0 - CLI with custom banner

This is an example CLI created with a custom banner.

Flags:

      --help   Get       help for the 'cli' command.
```
