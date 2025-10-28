package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Running stress test on %d cores...\n", numCPU)

	for i := 0; i < numCPU; i++ {
		go stresstestCPU(i)
	}

	select {}
}

func stresstestCPU(id int) {
	start := time.Now()
	x := 0.0001

	for {

		for i := 0; i < 10_000_000; i++ {
			x = math.Mod(x+math.Sin(float64(i))*0.1+0.5, 10.0)
		}

		if time.Since(start) > 5*time.Second {
			fmt.Printf("Goroutine %d: still alive, x = %.5f\n", id, x)
			start = time.Now()
		}
	}
}
