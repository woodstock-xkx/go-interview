package parallel

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"runtime"
	"testing"
)

func testParallelNumCounting(t *testing.T, N int) {
	dataLen := N
	dataRange := dataLen / 10
	data := make([]int, dataLen)
	for i, _ := range data {
		data[i] = int(rand.Int31n(int32(dataRange)))
	}

	serialNC := SerialNumCounting(data, dataRange)
	for numProcessors := 1; numProcessors <= runtime.NumCPU(); numProcessors++ {
		parallelNC := ParallelNumCounting(data, numProcessors, dataRange)
		assert.Equal(t, serialNC, parallelNC, "should equal")
	}
}

func TestParallelCounting(t *testing.T) {
	testCases := []int{1 << 20, 1 << 21, 1 << 22, 1 << 23, 1 << 24}
	for _, testCase := range testCases {
		testParallelNumCounting(t, testCase)
	}
}
