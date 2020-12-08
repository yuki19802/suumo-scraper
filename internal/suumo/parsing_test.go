package suumo

import "testing"

func TestExtractAgeYears(t *testing.T) {
	tests := []struct{
		input string
		expectedOutput int
	}{
		{
			input: "新築",
			expectedOutput: 0,
		},
		{
			input: "築16年",
			expectedOutput: 16,
		},
		{
			input: "築2年",
			expectedOutput: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			parsed, err := extractAgeYears(test.input)

			if err != nil {
				t.Fatal("Parsing failed:", err)
			}

			if parsed != test.expectedOutput {
				t.Fatalf("Got %d, want %d", parsed, test.expectedOutput)
			}
		})
	}
}

func TestExtractFloor(t *testing.T) {
	tests := []struct{
		input string
		expectedOutput int
	}{
		{
			input: "2階",
			expectedOutput: 2,
		},
		{
			input: "B1-1階",
			expectedOutput: 0,
		},
		{
			input: "-",
			expectedOutput: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			parsed, err := extractFloor(test.input)

			if err != nil {
				t.Fatal("Parsing failed:", err)
			}

			if parsed != test.expectedOutput {
				t.Fatalf("Got %d, want %d", parsed, test.expectedOutput)
			}
		})
	}
}
