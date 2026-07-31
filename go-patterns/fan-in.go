package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func fanIn(ctx context.Context, chans []chan int) chan int {
	outCh := make(chan int)
	wg := &sync.WaitGroup{}

	go func() {
		for _, ch := range chans {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case <-ctx.Done():
						fmt.Println("fanIn DONE")
						return
					case v, ok := <-ch:
						if !ok {
							fmt.Println("<-ch is closed.")
							return
						}
						outCh <- v
					}
				}
			}()
		}

		wg.Wait()
		println("after wg.WAIT() ===")
		close(outCh)
	}()

	return outCh
}

func runFanIn() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	chans := []chan int{ch1, ch2, ch3, ch4}
	ctx, cancel := context.WithTimeout(context.Background(), 3010*time.Millisecond)
	defer cancel()
	outCh := fanIn(ctx, chans)

	go func() {
		defer func() {
			for idx, ch := range chans {
				println(idx, " chan closed in main.")
				close(ch)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				println("LOG ===> ctx.Done()")
				return
			case <-time.After(1 * time.Second):
				for _, ch := range chans {
					ch <- int(time.Now().UnixMilli())
				}
			}

		}
	}()

Loop:
	for {
		v, ok := <-outCh
		if !ok {
			break Loop
		}
		println(" outCh ==> ", v)
	}

	println("END ==>")
}
