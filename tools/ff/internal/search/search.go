package search

import (
	"strings"
)

func SearchFiles(text string) ([]string, string, error) {
	stdout, stderr, err := run(text)
	if err != nil {
		return nil, stderr, err
	}

	lines := strings.Split(stdout, "\n")
	var nonEmpty []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmpty = append(nonEmpty, line)
		}
	}
	return nonEmpty, stderr, nil
}
