//multicopy is a multi-threaded URL retriever.  You provide
//a list of URLs to copy, multicopy does them as quickly as
//possible.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//var wg sync.WaitGroup
	//count := 10
	//c := make(chan string)
	//d := make(chan bool)
	b, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	urls := strings.Split(string(b), "\n")
	fmt.Println(urls)

}
