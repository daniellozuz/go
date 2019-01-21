package popcount

func PopCount(x uint64) int {
	total := 0
	for x > 0 {
		total += int(x & 1)
		x >>= 1
	}
	return total
}
