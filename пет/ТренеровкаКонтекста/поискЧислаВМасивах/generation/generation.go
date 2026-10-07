package generation

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerationArr(count, size int) [][]int {
	arrays := make([][]int, count)
	for i := 0; i < count; i++ {
		arrays[i] = make([]int, size)
		for j := 0; j < size; j++ {
			arrays[i][j] = rand.Intn(100)
		}
	}
	return arrays
}
