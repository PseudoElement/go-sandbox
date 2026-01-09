package iopack

import (
	"bufio"
	"log"
	"os"
	"time"
	"unicode"
	"unicode/utf8"
)

func ScanFileWords() error {
	file, err := os.Open("/Users/paveldavidovich/desktop/web/backend/go-sandbox/io/stream-data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(ScanNumbers)

	for scanner.Scan() {
		time.Sleep(50 * time.Millisecond)
		word := scanner.Text()
		println("Word ==> ", word)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func ScanNumbers(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip not numbers
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if unicode.IsNumber(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if unicode.IsSpace(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
