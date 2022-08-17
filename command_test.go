package ricecake

import (
	"reflect"
	"testing"
)

func TestCommand(t *testing.T) {
	t.Run("NewCommand", func(t *testing.T) {
		type test struct {
			testDescription string
			name            string
			description     string
			want            *Command
		}

		tests := []test{
			{
				testDescription: "basic command",
				name:            "Command",
				description:     "Test description",
				want: &Command{
					name:             "Command",
					shortDescription: "Test description",
					subCommandsMap:   make(map[string]*Command),
					hidden:           false,
				},
			},
			{
				testDescription: "basic command without description",
				name:            "Command",
				want: &Command{
					name:             "Command",
					shortDescription: "",
					subCommandsMap:   make(map[string]*Command),
					hidden:           false,
				},
			},
		}

		for _, tc := range tests {
			got := NewCommand(tc.name, tc.description)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("test failed (%s): got != want;\n got = %+v\n, want = %+v\n", tc.testDescription, got, tc.want)
			}
		}
	})

	t.Run("Run", func(t *testing.T) {
		cli := NewCLI("test", "Test CLI", "v0.1.0")
		cli.LongDescription("Long description for the test CLI.")
		// Should default to help menu.
		err := cli.Run()
		if err != nil {
			t.Errorf("test failed, err: %v", err)
		}

		// Set action.
		cli.Action(func() error {
			return nil
		})
		// Ensure that the action is being called.
		err = cli.Run("test")
		if err != nil {
			t.Errorf("test failed, err: %v", err)
		}

		// Check the help menu.
		err = cli.Run("--help")
		if err != nil {
			t.Errorf("test failed, err: %v", err)
		}

		sub := cli.NewSubCommand("sub", "Subcommand")
		sub.Action(func() error {
			return nil
		})

		// See help menu with newly-added subcommand.
		err = cli.Run("--help")
		if err != nil {
			t.Errorf("test failed, err: %v", err)
		}

		// Run subcommand.
		err = cli.Run("test", "sub")
		if err != nil {
			t.Errorf("test failed, err: %v", err)
		}
	})
}
