package iopack

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	ERR_TOO_MUCH_ARGS = errors.New("too much args")
)

func MainScan() {
	// path, err := exec.LookPath("node")
	// 	if err != nil {
	// 	panic(err)
	// }

	// arr := make([]string, 0, 10)
	// fmt.Println("Input fetched url:")
	// for len(arr) < 10 {
	// 	var value1 string
	// 	var value2 string
	// 	fmt.Scan(&value1, &value2)
	// 	arr = append(arr, value1, value2)
	// }

	scanner, err := NewEnvScan("BOOBOOK")
	if err != nil {
		panic(err)
	}

	var name string
	var company int32
	var event string

	n, err := scanner.Scan(&name, &company, &event)

	if errors.Is(err, ERR_TOO_MUCH_ARGS) {
		log.Printf("too much args passed: %v\n", err)
		return
	}

	log.Printf("successfully written %d vars: name - %v, company - %v, event - %v\n", n, name, company, event)

}

type EnvScan struct {
	envName string
	data    []string
}

func NewEnvScan(envName string) (*EnvScan, error) {
	godotenv.Load()

	value, ok := os.LookupEnv(envName)
	if !ok {
		return nil, fmt.Errorf("%s not set in .env file", envName)
	}

	scanner := &EnvScan{envName: envName}

	scanner.data = strings.Split(value, " ")

	return scanner, nil
}

/**
 * @param inputs are *pointers
 */
func (cs *EnvScan) Scan(inputs ...any) (n int, err error) {
	if len(inputs) > len(cs.data) {
		return -1, fmt.Errorf(
			"%w: expected - %d, got - %d",
			ERR_TOO_MUCH_ARGS,
			len(cs.data),
			len(inputs),
		)
	}

	var count int
	for idx, el := range cs.data {
		if idx < len(inputs) {
			variable, ok := inputs[idx].(*string)
			if !ok {
				return count, fmt.Errorf("argument %v is not a string pointer", variable)
			}

			*variable = el
			count++
		}
	}

	return count, nil
}
