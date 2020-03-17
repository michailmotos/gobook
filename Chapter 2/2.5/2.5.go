//2.5 performs a quick benchmark comparison between two implementations of counting set bits.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountRightClear returns the population count(number of set bits) of x
func PopCountRightClear(x uint64) int {
	var sum int
	for x != uint64(0) {
		x = x & (x - 1)
		sum++
	}
	return sum
}

// PopCount returns the population count(number of set bits) of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	var dataSet [10]uint64
	var sum int
	for i := 0; i < 10; i++ {
		dataSet[i] = rand.Uint64()
	}
	start := time.Now()
	for _, input := range dataSet {
		sum = PopCountRightClear(input)
		fmt.Printf("Number: %d , Set: %d\n", input, sum)
	}
	elapsed := time.Since(start)
	fmt.Printf("PopCountRightClear Elapsed time: %v\n", elapsed)

	start = time.Now()
	for _, input := range dataSet {
		sum = PopCount(input)
		fmt.Printf("Number: %d , Set: %d\n", input, sum)
	}
	elapsed = time.Since(start)
	fmt.Printf("PopCount Elapsed time: %v\n", elapsed)
}
