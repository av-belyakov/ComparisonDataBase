package commonlibs

import (
	"math"
)

// GetCountChunk выполняет деление с округлением в большую сторону
func GetCountChunk(maxSize int64, chunkSize int) int {
	ms := float64(maxSize)
	cs := float64(chunkSize)

	x := math.Floor(ms / cs)
	y := ms / cs

	if (y - x) != 0 {
		x++
	}

	return int(x)
}
