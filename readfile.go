package tetris

import (
	"log"
	"os"
)

// ReadFile reads a given file and returns a string slice split at carriage return, new line x 2
func ReadFile() string {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	return string(data)
}
