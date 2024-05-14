package equalsums

import (
	"sort"
)

func Solve(numbers []int, r shuffler) (subset1 []int, subset2 []int) {

	memo := map[int][]int{}

	n := len(numbers)
	k := n / 2
	// todo vary value of k

	for {
		r.Shuffle(n, func(i int, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
		sort.Ints(numbers[:k])
		s := sum(numbers[:k])
		// fmt.Println("sum of", numbers[:k], "is", s)
		if other, ok := memo[s]; ok {
			if !eq(numbers[:k], other) {
				// TODO remove common elements!
				return numbers[:k], other
			}
		} else {
			memo[s] = clone(numbers[:k])
		}
	}
}

// shuffler is implemented by *math/rand.Rand and *math/rand/v2.Rand
type shuffler interface {
	Shuffle(n int, swap func(i, j int))
}

func sum(a []int) int {
	s := 0
	for _, x := range a {
		s += x
	}
	return s
}

func clone(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
