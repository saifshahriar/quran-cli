package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func checkdbHealth() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	textPath := homeDir + "/.local/share/quran-cli/db/en.sahih.txt"
	f, err := os.Open(textPath)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	return string(content) // Return file content
}

func text() [][]string {
	text := checkdbHealth()
	readtext := csv.NewReader(strings.NewReader(text))
	readtext.Comma = '|'
	readtext.LazyQuotes = true
	lines, err := readtext.ReadAll()
	if err != nil {
		log.Fatalf("failed to read lines: %s", err)
	}
	return lines
}
