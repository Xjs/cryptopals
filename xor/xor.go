package xor

// Encrypt XOR-encrypts input with the given key
func Encrypt(input []byte, key []byte) []byte {
	result := make([]byte, len(input))

	for i, inputByte := range input {
		result[i] = inputByte ^ key[i%len(key)]
	}

	return result
}

// Single xors the input with a single byte
func Single(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i, inputByte := range input {
		result[i] = inputByte ^ key
	}
	return result
}
