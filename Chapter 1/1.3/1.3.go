//1.3 is an attempt to compare the string join func to the '+' operator
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[:], " "))
	elapsed := time.Since(start)
	fmt.Printf("strings.Join elapsed time: %v\n", elapsed)
	start = time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	elapsed = time.Since(start)
	fmt.Printf("Loop elapsed time: %v\n", elapsed)
}
