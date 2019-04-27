package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/Xjs/cryptopals/statistics"

	"github.com/Xjs/cryptopals/crack/english"

	"github.com/Xjs/cryptopals/xor"
)

func main() {
	input, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		log.Fatal(err)
	}

	// reference := getCharFrequency(englishText)

	scores := make(map[byte]statistics.Score)

	for i := 0; i < 256; i++ {
		b := byte(i)
		candidate := string(xor.Single(input, b))

		scores[b] = english.GetScore(candidate, english.ReferenceHistogram)
	}

	fs := statistics.NewByteScoreHistogram(scores)
	fmt.Println(fs)

	b := fs.GetHigh(0).Byte

	fmt.Println(string(xor.Single(input, b)))
}
