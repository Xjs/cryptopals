package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/Xjs/cryptopals/crack/english"
	"github.com/Xjs/cryptopals/sliceops"
	"github.com/Xjs/cryptopals/statistics"
)

func keysizeScore(input []byte, keysize int) (statistics.Score, error) {
	if len(input) < 2*keysize {
		return 0.0, errors.New("input too short")
	}

	first := input[:keysize]
	second := input[keysize : 2*keysize]

	return statistics.RelativeScore(sliceops.HammingDistance(first, second), keysize), nil
}

func main() {
	raw, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(raw)))
	n, err := base64.StdEncoding.Decode(decoded, raw)
	if err != nil {
		log.Fatal(err)
	}
	decoded = decoded[:n]
	log.Println("decoded", n, "bytes")

	scores := make(map[int]statistics.Score)
	for keysize := 2; keysize < 40; keysize++ {
		samples, _ := strconv.Atoi(os.Getenv("SAMPLES"))
		if samples == 0 {
			samples = 1
		}
		var sum statistics.Score
		for j := 0; j < samples; j++ {
			score, err := keysizeScore(decoded[2*j*keysize:], keysize)
			if err != nil {
				log.Println(err)
				continue
			}

			sum += score
		}
		sum /= statistics.Score(samples)
		scores[keysize] = sum
	}

	hist := statistics.NewSizeScoreHistogram(scores)
	log.Println(hist)

	var decrypted string
	var lowestNonEnglish int = 1000
	for _, tup := range hist {
		log.Println("Trying", tup.Size, "with score", tup.Score)

		result, nonEnglish, err := english.TryKeylen(decoded, tup.Size)
		if nonEnglish < lowestNonEnglish {
			lowestNonEnglish = nonEnglish
		}
		if err != nil {
			continue
		}

		if nonEnglish <= lowestNonEnglish && err == nil {
			log.Println("possible match")
			decrypted = string(result)
			break
		}
	}

	fmt.Println(decrypted)
}
