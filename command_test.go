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
}
