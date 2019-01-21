package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	total := 0
	for i := byte(0); i < 8; i++ {
		total += int(pc[byte(x>>(i*8))])
	}
	return total
}
