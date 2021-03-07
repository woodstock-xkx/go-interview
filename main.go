package main

import (
	"go-interview/parallel"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	dataLen := 1 << 28
	dataRange := 65536
	data := make([]int, dataLen)

	for i, _ := range data {
		data[i] = int(rand.Int31n(int32(dataRange)))
	}

	for i := 0; i < 4; i++ {

		serialStart := time.Now()
		parallel.SerialNumCounting(data, dataRange)
		serialTime := time.Since(serialStart)

		parallelStart := time.Now()
		//parallel.ParallelNumCounting(data, runtime.NumCPU())
		parallel.ParallelNumCounting(data, runtime.NumCPU(), dataRange)
		parallelTime := time.Since(parallelStart)

		println("serial: ", serialTime.String())
		println("parallel: ", parallelTime.String())

	}
}
