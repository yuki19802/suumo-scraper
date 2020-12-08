package suumo

import "testing"

func TestExtractAgeYears(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int
	}{
		{
			input:          "新築",
			expectedOutput: 0,
		},
		{
			input:          "築16年",
			expectedOutput: 16,
		},
		{
			input:          "築2年",
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
	tests := []struct {
		input          string
		expectedOutput int
	}{
		{
			input:          "2階",
			expectedOutput: 2,
		},
		{
			input:          "B1-1階",
			expectedOutput: 0,
		},
		{
			input:          "5-6階",
			expectedOutput: 0,
		},
		{
			input:          "-",
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

func TestExtractPrice(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int
	}{
		{
			input:          "9.8万円",
			expectedOutput: 98000,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			parsed, err := extractPriceYen(test.input)

			if err != nil {
				t.Fatal("Parsing failed:", err)
			}

			if parsed != test.expectedOutput {
				t.Fatalf("Got %d, want %d", parsed, test.expectedOutput)
			}
		})
	}
}

func TestExtractSquareMeters(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput float32
	}{
		{
			input:          "24.69m2",
			expectedOutput: 24.69,
		},
		{
			input:          "48.33m",
			expectedOutput: 48.33,
		},
		{
			input:          "22m2",
			expectedOutput: 22.00,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			parsed, err := extractSquareMeters(test.input)

			if err != nil {
				t.Fatal("Parsing failed:", err)
			}

			diff := parsed - test.expectedOutput

			if int(diff) != 0 {
				t.Fatalf("Got %.2f, want %.2f", parsed, test.expectedOutput)
			}
		})
	}
}
