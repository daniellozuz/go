package popcount

func PopCount(x uint64) int {
	total := 0
	for x > 0 {
		total++
		x = x & (x - 1)
	}
	return total
}
