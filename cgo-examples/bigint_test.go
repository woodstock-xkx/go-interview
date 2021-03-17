package cgo_examples

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func TestMult(t *testing.T) {
	N := 10000
	//N := 1
	testCases := []int{64, 128, 256, 512, 1024, 2048, 4096}

	for _, bits := range testCases {
		for i := 0; i < N; i++ {
			p, q := genPq(nil, bits)

			r1 := BnMult(p, q)
			r2 := new(big.Int).Mul(p, q)

			assert.Equal(t, r1, r2)
		}
	}
}

func genPq(rng *rand.Rand, bits int) (*big.Int, *big.Int) {
	max := new(big.Int)
	zero := new(big.Int)
	one := new(big.Int)

	max.SetBit(zero, bits+1, 1)
	max.Sub(max, one)

	if rng == nil {
		source := rand.NewSource(time.Now().UnixNano())
		rng = rand.New(source)
	}

	p := new(big.Int).Rand(rng, max)
	q := new(big.Int).Rand(rng, max)

	return p, q
}

func BenchmarkBnMul(b *testing.B) {
	p, q := genPq(nil, 1024)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		BnMult(p, q)
	}
	b.StopTimer()
}

func BenchmarkBigIntMul(b *testing.B) {
	p, q := genPq(nil, 1024)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		new(big.Int).Mul(p, q)
	}
	b.StopTimer()
}
