package suumo

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func extractAgeYears(raw string) (int, error) {
	if raw == "新築" {
		return 0, nil
	}

	trimmed := strings.TrimFunc(raw, func(r rune) bool {
		return !unicode.IsDigit(r)
	})

	return strconv.Atoi(trimmed)
}

func extractFloor(raw string) (int, error) {
	if strings.ContainsAny(raw, "B-") {
		return 0, nil
	}

	trimmed := strings.TrimFunc(raw, func(r rune) bool {
		return !unicode.IsDigit(r)
	})

	return strconv.Atoi(trimmed)
}

func extractPriceYen(raw string) (int, error) {
	if !strings.ContainsRune(raw, '万') {
		return 0, fmt.Errorf("string %q had unexpected amount", raw)
	}

	trimmed := strings.TrimFunc(raw, func(r rune) bool {
		return !unicode.IsDigit(r)
	})

	parsed, err := strconv.ParseFloat(trimmed, 32)

	if err != nil {
		return 0, fmt.Errorf("strconv.ParseFloat: %w", err)
	}

	return int(parsed*10000 + 0.1), nil
}

func extractSquareMeters(raw string) (float32, error) {
	trimmed := strings.TrimRight(raw, "2")
	trimmed = strings.TrimRight(trimmed, "m")

	parsed, err := strconv.ParseFloat(trimmed, 32)

	return float32(parsed), err
}
