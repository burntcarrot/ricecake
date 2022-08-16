package ricecake

import "fmt"

// defaultBanner returns a default banner for the CLI.
// Format: <name> <version> - <short description>
func defaultBanner(c *CLI) string {
	return fmt.Sprintf("%s %s - %s", c.Name(), c.Version(), c.ShortDescription())
}
