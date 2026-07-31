package main

import (
	"fmt"
	"time"
)

func rangeSimpleChan() {
	queue := make(chan string)

	go func() {
		queue <- "one"
		queue <- "two"
		queue <- "three"
		queue <- "last"
		time.Sleep(1 * time.Second)
		close(queue)
	}()

	for elem := range queue {
		fmt.Println("simple chan - ", elem)
	}
	println("===END_rangeSimpleChan===")
}

func rangeBufferedChan() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"

	go func() {
		queue <- "last"
		queue <- "last"
		queue <- "last"
		queue <- "last"
		queue <- "last"
		queue <- "last"
		queue <- "last"
		close(queue)
	}()

	for elem := range queue {
		fmt.Println("buff chan - ", elem)
	}
	println("===END_rangeBufferedChan===")
}
