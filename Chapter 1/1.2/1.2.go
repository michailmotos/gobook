//1.2 prints its command-line arguments & index.
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[:] {
		fmt.Printf("Index: %v \n    %s\n", idx, arg)
	}
}
