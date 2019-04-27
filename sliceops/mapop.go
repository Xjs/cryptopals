package sliceops

import "errors"

// A BinaryOperator is a binary operator on bytes
type BinaryOperator func(byte, byte) byte

// MapOperator maps an operator on two byte slices. It returns an error
// if the slices don't have equal lengths
func MapOperator(a, b []byte, op BinaryOperator) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("mapop: inputs are of different length")
	}

	result := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = op(a[i], b[i])
	}

	return result, nil
}

// XOR is bitwise XOR, or ^
var XOR BinaryOperator = func(a, b byte) byte { return a ^ b }
