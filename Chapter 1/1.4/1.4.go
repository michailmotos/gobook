// 1.4 prints the count and text of lines that appear more than once
// in the input files. It reads from stdin or from a list of named files.package dup2

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countFile := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countFile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countFile)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, countFile[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, countFile map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		countFile[input.Text()] = append(countFile[input.Text()], f.Name())
	}
	//NOTE: ignoring potential errors from input.Err()
}
