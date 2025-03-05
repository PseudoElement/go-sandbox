package iopack

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func Grep() {
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepErr, _ := grepCmd.StderrPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepErr.Close()
	grepIn.Close()

	grepBytes, _ := io.ReadAll(grepOut)
	errBytes, _ := io.ReadAll(grepErr)
	grepCmd.Wait()

	fmt.Println("grepBytes ==> ", string(grepBytes))
	fmt.Println("errBytes ==>", string(errBytes))
}

func ReadFileWords() error {
	// file, err := os.Open("/Users/paveldavidovich/desktop/web/backend/go-sandbox/io/stream-data.txt")
	file, err := os.Open("/Users/paveldavidovich/desktop/web/backend/go-sandbox/io/test.txt")
	defer file.Close()

	// fileContent, err := os.ReadFile("./stream-data.txt")
	buf := make([]byte, 20)

	if err != nil {
		panic(err)
	}

	for {
		// time.Sleep(500 * time.Millisecond)

		n, err := file.Read(buf)
		fmt.Println("Read ==> ", n)

		if err == io.EOF {
			break
		}
	}

	log.Println("Result buffer ==> ", string(buf))

	return nil
}
