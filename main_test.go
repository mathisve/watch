package main

import (
	"testing"
)

type testCase struct {
	args     []string
	argStart int
}

func TestParseArgs(t *testing.T) {
	var testCases = []testCase{
		{
			args:     []string{"/var/folders/r4/ljv4sv5x02q3vw9j1sb488cw0000gn/T/go-build2993121629/b001/exe/watch", "ls"},
			argStart: 1,
		},
		{
			args:     []string{"/var/folders/r4/ljv4sv5x02q3vw9j1sb488cw0000gn/T/go-build2993121629/b001/exe/watch", "-n", "10", "ls"},
			argStart: 3,
		},
		{
			args:     []string{"-n", "2", "echo", "hello"},
			argStart: 2,
		},
		{
			args:     []string{"-n", "echo", "hello"},
			argStart: 1,
		},
		{
			args:     []string{"hello", "123", "-n", "10", "ls"},
			argStart: 4,
		},
	}

	for _, test := range testCases {
		parsed := parseArgs(test.args)
		if parsed != test.argStart {
			t.Errorf("args: %+v, expected: %d, got: %d", test.args, test.argStart, parsed)
		}
	}
}
