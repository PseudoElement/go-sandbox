package main

func customMutex() {
	ch := make(chan struct{}, 1)
	done := make(chan struct{})
	idx := 0
	size := 10_000
	for range size {
		go func() {
			ch <- struct{}{}
			idx++
			<-ch
			if idx == size {
				done <- struct{}{}
			}
		}()
	}

	<-done
	println(idx)
}
