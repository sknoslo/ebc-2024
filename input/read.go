package input

import (
	"os"
	"strings"
)

func ReadNotes(path string) (string, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(buf)), nil
}
