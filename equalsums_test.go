package equalsums

import (
	"bytes"
	"math/rand/v2"
	"testing"
)

func testSolve(t testing.TB) {
	numbers := []int{738, 408, 588, 559, 976, 120, 713, 95, 620}
	var seed [32]byte
	copy(seed[:], bytes.Repeat(nil, 42))
	r := rand.New(rand.NewChaCha8(seed))
	sub1, sub2 := Solve(numbers, r)

	if eq(sub1, sub2) {
		t.Errorf("Identical subsets %v, %v", sub1, sub2)
	}

	if sum(sub1) != sum(sub2) {
		t.Errorf("Wrong solution %v, %v", sub1, sub2)
	}
}

func TestSolve(t *testing.T) {
	testSolve(t)
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testSolve(b)
	}
}
