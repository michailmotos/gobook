// Fetchall fetches URLs in parallel and reports their times and sizes.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	f, err := os.Create("/tmp/dump.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	toWrite := fmt.Sprintf("%f", time.Since(start).Seconds())
	toWrite = toWrite + "s elapsed\n"
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	_, err = f.WriteString(toWrite)
	if err != nil {
		fmt.Println(err)
		f.Close()
		os.Exit(1)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
