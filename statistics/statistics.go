package statistics

import (
	"fmt"
	"sort"
)

// A Score is a number that is used to rank something.
type Score float64

// RelativeScore returns numerator/denominator as floating point Score
func RelativeScore(numerator, denominator int) Score {
	return Score(numerator) / Score(denominator)
}

// A ByteScoreHistogramEntry is a single entry of a ByteScoreHistogram
type ByteScoreHistogramEntry struct {
	Byte  byte
	Score Score
}

// A ByteScoreHistogram is a map between bytes and counts. It can be sorted.
type ByteScoreHistogram []ByteScoreHistogramEntry

func (bh ByteScoreHistogram) Less(a, b int) bool { return bh[a].Score < bh[b].Score }
func (bh ByteScoreHistogram) Swap(a, b int)      { bh[a], bh[b] = bh[b], bh[a] }
func (bh ByteScoreHistogram) Len() int           { return len(bh) }

// NewByteScoreHistogram creates a sorted ByteScoreHistogram from a map byte -> int
func NewByteScoreHistogram(m map[byte]Score) ByteScoreHistogram {
	var result ByteScoreHistogram
	for b, score := range m {
		result = append(result, ByteScoreHistogramEntry{Byte: b, Score: score})
	}
	sort.Sort(result)
	return result
}

// GetHigh gets the index-th highest entry from the histogram
func (bh ByteScoreHistogram) GetHigh(index int) ByteScoreHistogramEntry {
	sort.Sort(bh)
	return bh[len(bh)-1-index]
}

func (bh ByteScoreHistogram) String() string {
	var result string
	for i, entry := range bh {
		result += fmt.Sprintf("%q: %.2f", entry.Byte, entry.Score)
		if i != len(bh)-1 {
			result += ", "
		}
	}
	return result
}

// A SizeScoreHistogramEntry is a single entry of a SizeScoreHistogram
type SizeScoreHistogramEntry struct {
	Size  int
	Score Score
}

// A SizeScoreHistogram is a map between sizes and counts. It can be sorted.
type SizeScoreHistogram []SizeScoreHistogramEntry

func (bh SizeScoreHistogram) Less(a, b int) bool { return bh[a].Score < bh[b].Score }
func (bh SizeScoreHistogram) Swap(a, b int)      { bh[a], bh[b] = bh[b], bh[a] }
func (bh SizeScoreHistogram) Len() int           { return len(bh) }

// NewSizeScoreHistogram creates a sorted SizeScoreHistogram from a map size -> Score
func NewSizeScoreHistogram(m map[int]Score) SizeScoreHistogram {
	var result SizeScoreHistogram
	for b, score := range m {
		result = append(result, SizeScoreHistogramEntry{Size: b, Score: score})
	}
	sort.Sort(result)
	return result
}

// GetHigh gets the index-th highest entry from the histogram
func (bh SizeScoreHistogram) GetHigh(index int) SizeScoreHistogramEntry {
	sort.Sort(bh)
	return bh[len(bh)-1-index]
}

func (bh SizeScoreHistogram) String() string {
	var result string
	for i, entry := range bh {
		result += fmt.Sprintf("%d: %.2f", entry.Size, entry.Score)
		if i != len(bh)-1 {
			result += ", "
		}
	}
	return result
}

// A ByteHistogramEntry is a single entry of a ByteHistogram
type ByteHistogramEntry struct {
	Byte  byte
	Count int
}

// A ByteHistogram is a map between bytes and counts. It can be sorted.
type ByteHistogram []ByteHistogramEntry

func (bh ByteHistogram) Less(a, b int) bool { return bh[a].Count < bh[b].Count }
func (bh ByteHistogram) Swap(a, b int)      { bh[a], bh[b] = bh[b], bh[a] }
func (bh ByteHistogram) Len() int           { return len(bh) }

// NewByteHistogram creates a sorted ByteHistogram from a map byte -> int
func NewByteHistogram(m map[byte]int) ByteHistogram {
	var result ByteHistogram
	for b, count := range m {
		result = append(result, ByteHistogramEntry{Byte: b, Count: count})
	}
	sort.Sort(result)
	return result
}

// GetHigh gets the index-th highest entry from the histogram
func (bh ByteHistogram) GetHigh(index int) ByteHistogramEntry {
	sort.Sort(bh)
	return bh[len(bh)-1-index]
}

func (bh ByteHistogram) String() string {
	var result string
	for i, entry := range bh {
		result += fmt.Sprintf("%q: %d", entry.Byte, entry.Count)
		if i != len(bh)-1 {
			result += ", "
		}
	}
	return result
}

// NewByteHistogramFromRunes creates a ByteHistogram by converting runes to bytes
func NewByteHistogramFromRunes(rh map[rune]int) ByteHistogram {
	var bh ByteHistogram
	for b, c := range rh {
		bh = append(bh, ByteHistogramEntry{byte(b), c})
	}
	sort.Sort(bh)
	return bh
}
