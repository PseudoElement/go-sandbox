package main

import (
	"fmt"
	"time"
)

func rangeSimpleChan() {
	queue := make(chan string)

	go func() {
		queue <- "one"
		time.Sleep(1 * time.Second)
		queue <- "two"
		time.Sleep(1 * time.Second)
		queue <- "three"
		time.Sleep(1 * time.Second)
		queue <- "last"
		close(queue)
	}()

	for elem := range queue {
		fmt.Println("simple chan - ", elem)
	}
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
}
