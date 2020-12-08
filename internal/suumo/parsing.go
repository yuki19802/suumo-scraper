package suumo

import (
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
	if strings.HasPrefix(raw, "B") || raw == "-" {
		return 0, nil
	}


	trimmed := strings.TrimFunc(raw, func(r rune) bool {
		return !unicode.IsDigit(r)
	})

	return strconv.Atoi(trimmed)
}
