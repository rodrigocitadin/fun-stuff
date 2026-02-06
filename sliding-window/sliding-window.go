package slidingwindow

import (
	"math"
)

func SlidingWindow(str string, occurrences int) int {
	runes := []rune(str)
	counter := make(map[rune]int)
	counter[runes[0]] = 1

	l, r := 0, 0
	max := 1

	for r < len(runes)-1 {
		r += 1
		counter[runes[r]] += 1

		for counter[runes[r]] == occurrences {
			counter[runes[l]] -= 1
			l += 1
		}

		max = int(math.Max(
			float64(max),
			float64(r-l+1)))
	}

	return max
}
