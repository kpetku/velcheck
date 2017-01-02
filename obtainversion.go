package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func obtainVersion(input string) (string, error) {
	if file, err := os.Open(input); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "<version>") {
				return strings.Fields(strings.TrimRight(strings.TrimLeft(strings.TrimSpace(scanner.Text()), "<version>"), "</version>"))[0], nil
			}
		}
		if err = scanner.Err(); err != nil {
			return "", err
		}
	}
	return "", errors.New("Missing version information")
}
