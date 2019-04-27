package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/Xjs/cryptopals/sliceops"
)

func main() {
	input, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		log.Fatal(err)
	}
	mask, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		log.Fatal(err)
	}

	result, err := sliceops.MapOperator(input, mask, sliceops.XOR)
	if err != nil {
		log.Fatal(err)
	}

	output := hex.EncodeToString(result)

	if output == "746865206b696420646f6e277420706c6179" {
		fmt.Println("exercise passed")
	}
}
