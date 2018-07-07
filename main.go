//multicopy is a multi-threaded URL retriever.  You provide
//a list of URLs to copy, multicopy does them as quickly as
//possible.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func run(i int, c chan string, d chan bool) {
	for {
		select {
		case u := <-c:
			fmt.Printf("%d: '%s'\n", i, u)
		case <-d:
			fmt.Printf("%d: done\n", i)
			return
		default:
		}
	}
}
func main() {
	b, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	var wg sync.WaitGroup
	count := 10
	c := make(chan string)
	d := make(chan bool)
	for i := 1; i <= count; i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			run(i, c, d)
		}(i)
	}
	urls := strings.Split(string(b), "\n")
	for _, u := range urls {
		c <- u
	}
	close(d)
	wg.Wait()
}
