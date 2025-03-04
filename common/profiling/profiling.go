package profiling

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

func Run() {
	cpuProfileFile, err := os.Create("./prof/cpu.prof")
	if err != nil {
		panic(err)
	}
	defer cpuProfileFile.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	fmt.Println("CPU profile written to /prof/cpu.prof")

	memProfileFile, err := os.Create("./prof/mem.prof")
	if err != nil {
		panic(err)
	}
	defer memProfileFile.Close()

	// Write memory profile to file
	if err := pprof.WriteHeapProfile(memProfileFile); err != nil {
		panic(err)
	}

	fmt.Println("Memory profile written to /prof/mem.prof")

	// Start tracing
	traceFile, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer traceFile.Close()

	if err := trace.Start(traceFile); err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("Tracing written to /prof/trace.out")
}

func heavyCalc() {
	for i := 0; i < 1_000; i++ {
		time.Sleep(10 * time.Millisecond)
		randNum := rand.Intn(100)
		square := math.Pow(float64(randNum), 2)
		if i%100 == 0 {
			fmt.Printf("%v ** 2 = %v\n", randNum, square)
		}
	}
}
