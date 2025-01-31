package funcs

import (
	"io"
	"os"
)

func WriteFile(content string) {
	err := os.WriteFile("./written-file.csv", []byte(""), 0644)
	file, err := os.OpenFile("./written-file.csv", os.O_WRONLY, os.ModeAppend)
	_, err = io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
}
