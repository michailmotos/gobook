//2.4 performs a quick benchmark comparison between two implementations of counting set bits.
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

// PopCountShift returns the population count(number of set bits) of x
func PopCountShift(x uint64) int {
	var comparator uint64 = 1
	var sum int
	for i := 0; i < 64; i++ {
		if x&comparator == 1 {
			sum++
		}
		x >>= 1
	}
	return sum
}

//PopCountLoop return the population count(number of set bits) of x
func PopCountLoop(x uint64) int {
	var sum int
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

func main() {
	var dataSet [10]uint64
	var sum int
	for i := 0; i < 10; i++ {
		dataSet[i] = rand.Uint64()
	}
	start := time.Now()
	for _, input := range dataSet {
		sum = PopCountShift(input)
		fmt.Printf("Number: %d , Set: %d\n", input, sum)
	}
	elapsed := time.Since(start)
	fmt.Printf("PopCount Elapsed time: %v\n", elapsed)

	start = time.Now()
	for _, input := range dataSet {
		sum = PopCountLoop(input)
		fmt.Printf("Number: %d , Set: %d\n", input, sum)
	}
	elapsed = time.Since(start)
	fmt.Printf("PopCountLoop Elapsed time: %v\n", elapsed)
}
