//2.2 converts its numeric argument or std input to length and meters
package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	var input float64
	if len(os.Args) <= 1 {
		fmt.Print("Enter temperature reading to convert: ")
		fmt.Scanln(&input)
		f := tempconv.Fahrenheit(input)
		c := tempconv.Celsius(input)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		}
	}
}
