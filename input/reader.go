package input

import (
	"os"
	"strings"
)

/*
Splits the input string into lines, handling both Unix and Windows line endings.
*/
func LoadInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		println("Error reading file:", err)
		return []string{}
	}

	// Normalize Windows and Unix line endings and trim trailing newline
	s := strings.ReplaceAll(string(data), "\r\n", "\n")
	s = strings.TrimRight(s, "\n")

	if s == "" {
		return []string{}
	}
	return strings.Split(s, "\n")
}
