package sliceops

import (
	"math/bits"
)

// HammingDistance calculates the number of differing bits between the two input slices.
func HammingDistance(a, b []byte) int {
	l := len(a)
	if len(b) > l {
		l = len(b)
	}

	normalisedA := make([]byte, l)
	normalisedB := make([]byte, l)

	copy(normalisedA, a)
	copy(normalisedB, b)

	var result int

	for i := 0; i < l; i++ {
		result += bits.OnesCount8(uint8(normalisedA[i] ^ normalisedB[i]))
	}

	return result
}
