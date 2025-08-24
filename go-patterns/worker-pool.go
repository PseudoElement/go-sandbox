package main

import (
	"context"
	"log"
	"math"
	"sync"
	"time"
)

type WorkerResp struct {
	Res      int
	WorkerId int
}

func runWorkerPool() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan int)
	callback := func(num int) int {
		return int(math.Pow(float64(num), 2))
	}

	outCh := createWorkers(ch, callback, 5, ctx)

	go func() {
		for val := range 20 {
			ch <- val
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	for resp := range outCh {
		log.Printf("response ==>%+v \n", resp)
	}
}

func createWorkers(ch <-chan int, callback func(num int) int, workersCount int, ctx context.Context) chan WorkerResp {
	outChan := make(chan WorkerResp)
	wg := &sync.WaitGroup{}
	wg.Add(workersCount)

	go func() {
		wg.Wait()
		close(outChan)
	}()

	for idx := range workersCount {
		go func() {
		Loop:
			for {
				select {
				case <-ctx.Done():
					println("timeout!", idx)
					wg.Done()
					break Loop
				case val, ok := <-ch:
					if !ok {
						wg.Done()
						break Loop
					} else {
						outChan <- WorkerResp{Res: callback(val), WorkerId: idx}
					}
				}
			}
		}()
	}

	return outChan
}
