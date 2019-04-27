package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Xjs/cryptopals/crack/english"
)

func main() {
	text, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(text), "\n")

	for _, line := range lines {
		input, _ := hex.DecodeString(line)
		result, key, score, err := english.FindSingleByteKey(input)
		if err != nil {
			continue
		}
		log.Printf("key: %q (score: %.2f)\n", key, score)
		fmt.Printf(result)
	}
}
