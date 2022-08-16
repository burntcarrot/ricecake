<div align = "center">
    <img src="static/ricecake.png">
</div>

# ricecake

`ricecake` is a lightweight framework to build CLI applications in Go.

## Features

- Zero dependencies <sup>only uses <a href="https://github.com/spf13/pflag">pflag</a>; based on the standard library</sup>
- Support for subcommands
- POSIX-compliant flags
- Action-based API (inspired by [urfave/cli](https://github.com/urfave/cli))
- Init Hook (hook for customizing behavior during CLI initialization)
- Auto-generated help menu!
- Custom banners
- and more!

## Installation

Installing `ricecake` is easy, just import:

```go
import "github.com/burntcarrot/ricecake"
```

## Usage

A list of examples:

- [Create a basic CLI with ricecake](./examples/cli/main.go)
- coming soon!

## License

`ricecake` is released under the MIT license. See [LICENSE](./LICENSE).
