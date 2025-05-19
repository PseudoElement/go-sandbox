package utils

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Str struct {
	Data int
}

func UseAtomicPointerForStruct(itersCount int) {
	var ptr atomic.Pointer[Str]

	ptr.Store(&Str{Data: 1}) // Initial value

	var wg sync.WaitGroup

	for range itersCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				current := ptr.Load()
				newVal := &Str{Data: current.Data + 1}
				// Try to swap in the new value
				if ptr.CompareAndSwap(current, newVal) {
					break // Success
				}
				// Else, retry (another goroutine beat us to it)
			}
		}()
	}

	wg.Wait()

	fmt.Println("atomic pointer res - ", ptr.Load().Data)
}

func UseMutexForStruct(itersCount int) {
	var ptr = &Str{Data: 1}

	var wg sync.WaitGroup
	var mu sync.RWMutex
	for range itersCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			ptr.Data++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("mutex res - ", ptr.Data)
}
