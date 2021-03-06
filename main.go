package main

import (
	"go-interview/parallel"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	dataLen := 1 << 28
	data := make([]int, dataLen)

	for i, _ := range data {
		data[i] = int(rand.Int31n(int32(65536)))
	}

	for i := 0; i < 4; i++ {

		serialStart := time.Now()
		parallel.SerialNumCounting(data)
		serialTime := time.Since(serialStart)

		parallelStart := time.Now()
		//parallel.ParallelNumCounting(data, runtime.NumCPU())
		parallel.ParallelNumCounting(data, runtime.NumCPU())
		parallelTime := time.Since(parallelStart)

		println("serial: ", serialTime.String())
		println("parallel: ", parallelTime.String())

	}
}
