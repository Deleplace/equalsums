package equalsums

import (
	"bytes"
	"math/rand"
	randv2 "math/rand/v2"
	"testing"
)

func testSolve(t testing.TB, r shuffler) {
	numbers := []int{738, 408, 588, 559, 976, 120, 713, 95, 620}
	sub1, sub2 := Solve(numbers, r)

	if eq(sub1, sub2) {
		t.Errorf("Identical subsets %v, %v", sub1, sub2)
	}

	if sum(sub1) != sum(sub2) {
		t.Errorf("Wrong solution %v, %v", sub1, sub2)
	}
}

func TestSolve(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	testSolve(t, r)
}

func BenchmarkSolveRand(b *testing.B) {
	for range b.N {
		r := rand.New(rand.NewSource(42))
		testSolve(b, r)
	}
}

func BenchmarkSolveChaCha8(b *testing.B) {
	for range b.N {
		var seed [32]byte
		copy(seed[:], bytes.Repeat([]byte{42}, 32))
		r := randv2.New(randv2.NewChaCha8(seed))
		testSolve(b, r)
	}
}

func BenchmarkSolvePCG(b *testing.B) {
	for range b.N {
		var seed [32]byte
		copy(seed[:], bytes.Repeat([]byte{42}, 32))
		r := randv2.New(randv2.NewPCG(42, 42))
		testSolve(b, r)
	}
}
