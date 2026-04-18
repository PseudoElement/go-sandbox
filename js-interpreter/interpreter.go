// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"log"
	"regexp"
	"strings"
)

func interpret(input string) {
	if strings.HasPrefix(input, "console.log") && strings.HasSuffix(input, "\");") {
		rest := strings.TrimPrefix(input, `console.log("`)
		rest = strings.TrimSuffix(rest, `");`)

		re := regexp.MustCompile(`("|'),\s*?("|')`)
		splitted := re.Split(rest, -1)

		var args []any
		args = append(args, "[LOG] ")
		for _, el := range splitted {
			args = append(args, el)
		}

		log.Println(args...)
	} else {
		panic("Type error: Unknown input type.")
	}
}

func main() {
	interpret(`console.log("Hello,", "world!");`)
	interpret(`console.log("Hello,", "world!", "Kuzya is CTO.");`)
}
