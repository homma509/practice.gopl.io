package main

var pc [256]byte // 8bit * 32

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountOneLiner(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountTableLoop(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(uint(i)*8))])
	}
	return count
}
