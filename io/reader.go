package iopack

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type DotsReader struct {
	text string
}

func NewDotsReader(text string) *DotsReader {
	return &DotsReader{text}
}

func (r *DotsReader) Read(b []byte) (n int, err error) {
	var count int
	for i, char := range r.text {
		if i > len(b) {
			return count, io.EOF
		}
		if string(char) == "." {
			b[count] = r.text[i]
			count++
		}
	}

	return count, io.EOF
}

func ReadBuffer() error {
	emptyBuf := make([]byte, 1000)
	fullBuf := []byte("Length is 10.")

	var reader = bytes.NewReader(fullBuf)

	// newBuf, err := io.ReadAll(reader)
	// _, err := io.ReadFull(reader, emptyBuf)
	_, err := reader.Read(emptyBuf)

	log.Println("ReadBuffer_Empty ==> ", emptyBuf)
	log.Println("ReadBuffer_Full ==> ", string(fullBuf))
	log.Println("ReadBuffer_Error ==> ", err)

	return err
}

func Grep() {
	grepCmd := exec.Command("grep", "-E", "good|bimba")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepErr, _ := grepCmd.StderrPipe()

	grepCmd.Start()
	grepIn.Write([]byte("bimba grep ola amigo\ngoodbye grep\n"))
	grepIn.Write([]byte("you suck\ngood morning"))

	grepIn.Close()
	defer grepErr.Close()
	defer grepOut.Close()

	grepBytes, _ := io.ReadAll(grepOut)
	errBytes, _ := io.ReadAll(grepErr)
	grepCmd.Wait()

	fmt.Println("grepBytes ==>\n", string(grepBytes))
	fmt.Println("errBytes ==>", string(errBytes))
}

func ReadFileWords() error {
	file, err := os.Open("/Users/paveldavidovich/desktop/web/backend/go-sandbox/io/stream-data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// fileContent, err := os.ReadFile("./stream-data.txt")
	// fileContent, err := os.ReadFile("./stream-data.txt")
	buf := make([]byte, 20)

	for {
		n, err := file.Read(buf)
		fmt.Println("Read ==> ", n)

		if err == io.EOF {
			break
		}
	}

	log.Println("Result buffer ==> ", string(buf))

	return nil
}
