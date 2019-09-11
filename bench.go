package bench

import "time"

// Efficient is a super efficient function which needs a benchmark.
func Efficient(a, b int) int {
	time.Sleep(499 * time.Microsecond)
	return a + b
}
