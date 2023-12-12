package board

import (
	"math/rand"
	"time"
)

func generateRandomNumber(n int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(n)
}

func processLine(original []int, headToTail bool) []int {
	processed := make([]int, 0, len(original))

	if headToTail {
		last := 0
		hasLast := false
		for i := len(original) - 1; i >= 0; i-- {
			if original[i] == 0 {
				continue
			}

			if last == 0 {
				last = original[i]
				hasLast = true
				continue
			}

			if original[i] == last {
				hasLast = false
				processed = append(processed, last*2)
				last = 0
			} else {
				processed = append(processed, last)
				last = original[i]
				hasLast = true
			}

		}

		if hasLast {
			processed = append(processed, last)
		}
	} else {
		last := 0
		hasLast := false
		for i := 0; i < len(original); i++ {
			if original[i] == 0 {
				continue
			}

			if last == 0 {
				last = original[i]
				hasLast = true
				continue
			}

			if original[i] == last {
				hasLast = false
				processed = append(processed, last*2)
				last = 0
			} else {
				processed = append(processed, last)
				last = original[i]
				hasLast = true
			}

		}

		if hasLast {
			processed = append(processed, last)
		}
	}

	res := make([]int, len(original))
	offset := 0

	if headToTail {
		offset = len(original) - 1
		for i := 0; i < len(processed); i++ {
			res[offset] = processed[i]
			offset -= 1
		}

	} else {
		offset = 0
		for i := 0; i < len(processed); i++ {
			res[offset] = processed[i]
			offset += 1
		}
	}

	return res

}
