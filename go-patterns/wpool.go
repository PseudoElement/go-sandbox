package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func runWPool() {
	wpool := NewWorkerPool(10)
	logsChan := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for idx := range 100 {
			time.Sleep(1 * time.Millisecond)
			logsChan <- fmt.Sprintf("Log number %d", idx)
		}
		close(logsChan)
	}()
	go func() {
		defer wpool.Close()
		for newLog := range logsChan {
			wpool.DoWork(ctx, func() any {
				time.Sleep(1000 * time.Millisecond)
				log.Println("worker called log function: ", newLog)
				return "success"
			})
		}
	}()
	go func() {
		defer wg.Done()
		for resp := range wpool.Results() {
			log.Println("result - ", resp)
		}
	}()

	wg.Wait()

	_, ok1 := <-wpool.results
	_, ok2 := <-wpool.workers

	log.Println(ok1, ok2)
}

type WorkerPool struct {
	workers   chan func() any
	results   chan any
	jobsCount int
	wg        *sync.WaitGroup
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workers:   make(chan func() any, workerCount),
		results:   make(chan any),
		jobsCount: 0,
		wg:        &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) DoWork(ctx context.Context, fn func() any) {
	wp.wg.Add(1)
	select {
	case <-ctx.Done():
		wp.wg.Done()
		return
	case wp.workers <- fn:
		go func() {
			defer func() {
				wp.wg.Done()
				<-wp.workers
			}()
			select {
			case <-ctx.Done():
				println("timeout")
				return
			default:
				res := fn()
				wp.results <- res
				return
			}
		}()
	}
}

func (wp *WorkerPool) Close() {
	go func() {
		wp.wg.Wait()
		close(wp.workers)
		close(wp.results)
	}()
}

func (wp *WorkerPool) WaitForEnd() {
	for range wp.Results() {
	}
}

func (wp *WorkerPool) Results() <-chan any {
	return wp.results
}
