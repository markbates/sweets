package sweets

import (
	"strings"
	"unicode"
)

func TrimLeftSpace(s string) string {
	var lines []string

	for _, line := range strings.Split(s, "\n") {
		if len(line) == 0 {
			lines = append(lines, line)
			continue
		}

		rns := []rune(line)
		if !unicode.IsSpace(rns[0]) {
			return s
		}
		lines = append(lines, string(rns[1:]))
	}

	s = strings.Join(lines, "\n")

	if len(strings.TrimSpace(s)) == 0 {
		return s
	}

	return TrimLeftSpace(s)
}
