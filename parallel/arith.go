package parallel

import (
	"fmt"
	"time"
)

func SerialNumCounting(data []int, numRange int) map[int]uint {
	counts := make(map[int]uint, numRange)
	for _, d := range data {
		counts[d]++
	}
	return counts
}

func ParallelNumCounting(data []int, numProcessors int, numRange int) map[int]uint {
	countsCh := make(chan map[int]uint, numProcessors)
	n := len(data)
	if n < numProcessors {
		return SerialNumCounting(data, numRange)
	}
	for id := 0; id < numProcessors; id++ {
		chunkSize := (n + numProcessors - 1) / numProcessors
		start, end := chunkSize*id, chunkSize*(id+1)
		if end > n {
			end = n
		}

		go func(id int, dataSlice []int) {
			startTime := time.Now()
			countsCh <- SerialNumCounting(dataSlice, numRange)
			fmt.Printf("goroutine[%d] takes {%s} to count data of length %d\n",
				id, time.Since(startTime).String(), len(dataSlice))
		}(id, data[start:end])
	}

	finalCounts := make(map[int]uint, numRange)
	for i := 0; i < numProcessors; i++ {
		counts := <-countsCh
		for k, v := range counts {
			finalCounts[k] += v
		}
	}
	close(countsCh)
	return finalCounts
}
