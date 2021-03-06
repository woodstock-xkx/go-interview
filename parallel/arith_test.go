package parallel

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"runtime"
	"testing"
)

func testParallelNumCounting(t *testing.T, N int) {
	dataLen := N
	data := make([]int, dataLen)
	for i, _ := range data {
		data[i] = int(rand.Int31n(int32(dataLen / 10)))
	}

	serialNC := SerialNumCounting(data)
	for numProcessors := 1; numProcessors <= runtime.NumCPU(); numProcessors++ {
		parallelNC := ParallelNumCounting(data, numProcessors)
		assert.Equal(t, serialNC, parallelNC, "should equal")
	}
}

func TestParallelCounting(t *testing.T) {
	testCases := []int{1 << 20, 1 << 21, 1 << 22, 1 << 23, 1 << 24}
	for _, testCase := range testCases {
		testParallelNumCounting(t, testCase)
	}
}
