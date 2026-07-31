package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"sync"
	"time"
)

func fanOut() {
	chOut := sendQueries(
		[]string{"localhost:8080", "146.124.34.19:8081", "ss", "sa", "213"},
		"SELECT * FROM users WHERE name IN ('Alice', 'Andersen')",
	)
	for val := range chOut {
		log.Println("RESULT: ", val)
	}
	batchReqs([]string{"localhost:8080", "146.124.34.19:8081"}, []byte(""))
}

func makeQuery(shardAddr string, query string) string {
	delay := 1 + rand.IntN(5)
	time.Sleep(time.Duration(int(delay) * int(time.Second)))
	return fmt.Sprintf("[%s] result_%d", shardAddr, delay)
}

func sendQueries(addresses []string, query string) <-chan string {
	chOut := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(len(addresses))

	go func() {
		wg.Wait()
		close(chOut)
	}()
	for _, addr := range addresses {
		go func() {
			chOut <- makeQuery(addr, query)
			wg.Done()
		}()
	}
	return chOut
}
