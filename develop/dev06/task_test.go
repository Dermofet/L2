package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		input       string
		fields      string
		delimiter   string
		separated   bool
		expectedOut string
	}{
		// Test selecting specific fields
		{
			input:       "apple\tbanana\torange\napple\tbanana\torange\n",
			fields:      "1,3",
			delimiter:   "\t",
			separated:   false,
			expectedOut: "apple\torange\napple\torange\n",
		},
		// Test using a different delimiter
		{
			input:       "apple,banana,orange\napple,banana,orange\n",
			fields:      "1,3",
			delimiter:   ",",
			separated:   false,
			expectedOut: "apple,orange\napple,orange\n",
		},
		// Test only output lines containing delimiter
		{
			input:       "apple,banana\napple\n",
			fields:      "1,2",
			delimiter:   ",",
			separated:   true,
			expectedOut: "apple,banana\n",
		},
	}

	for _, test := range tests {
		in := strings.NewReader(test.input)
		out := &bytes.Buffer{}

		args := &InputArgs{
			Fields:    test.fields,
			Delimiter: test.delimiter,
			Separated: test.separated,
		}

		Cut(in, out, args)

		result := out.String()

		if result != test.expectedOut {
			t.Errorf("For input '%s' with fields '%s' and delimiter '%s' (separated=%v), expected output '%s', but got '%s'",
				test.input, test.fields, test.delimiter, test.separated, test.expectedOut, result)
		}
	}
}
