package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		inputArgs   *InputArgs
		inputFile   string
		expectedOut string
	}{
		// Test basic match
		{
			inputArgs: &InputArgs{
				Pattern: "Hello",
				Filenames: []string{
					"testfile.txt",
				},
			},
			expectedOut: "Hello, world!\n",
		},
		// // Test ignore case match
		// {
		// 	inputArgs: &InputArgs{
		// 		Pattern:    "WORLD",
		// 		IgnoreCase: true,
		// 		Filenames: []string{
		// 			"testfile.txt",
		// 		},
		// 	},
		// 	expectedOut: "Hello, world!\n",
		// },
		// // Test invert match
		// {
		// 	inputArgs: &InputArgs{
		// 		Pattern: "world",
		// 		Invert:  true,
		// 		Filenames: []string{
		// 			"testfile.txt",
		// 		},
		// 	},
		// 	expectedOut: "Bu-Bu-Bu\n",
		// },
		// // Test fixed match
		// {
		// 	inputArgs: &InputArgs{
		// 		Pattern: "Hello, world!",
		// 		Fixed:   true,
		// 		Filenames: []string{
		// 			"testfile.txt",
		// 		},
		// 	},
		// 	expectedOut: "Hello, world!\n",
		// },
	}

	for _, test := range tests {
		old := os.Stdout // keep backup of the real stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		Grep(test.inputArgs)

		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)
		os.Stdout = old // restoring the real stdout
		out := buf.String()

		if out != test.expectedOut {
			t.Errorf("Expected output %s, but got %s", test.expectedOut, out)
		}
	}
}
