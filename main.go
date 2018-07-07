//multicopy is a multi-threaded URL retriever.  You provide
//a list of URLs to copy, multicopy does them as quickly as
//possible.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
)

//Run reads URLs from a channel and writes them to disk.
//Errors are logged and are not fatal.  Processing continues
//until the done channel has contents or is closed.
func Run(c chan string, dir string, done chan bool) {
	var errLog = log.New(os.Stderr, "", log.LstdFlags)
	for {
		select {
		case u := <-c:
			_, err := Store(u, dir)
			if err != nil {
				errLog.Printf("Error: %v %s\n", err, u)
			} else {
				log.Printf("%s\n", u)
			}
		case <-done:
			return
		default:
		}
	}
}

//Store saves a URL to disk.  The contents are stored
//in the provided directory keeping the URL's directory
//structure intact.
func Store(link string, dir string) (int64, error) {
	r, err := http.Get(link)
	if err != nil {
		return 0, err
	}
	if r.StatusCode != 200 {
		m := fmt.Sprintf("%v (%v)", r.StatusCode, http.StatusText(r.StatusCode))
		return 0, errors.New(m)
	}
	defer r.Body.Close()

	u, err := url.Parse(link)
	if err != nil {
		return 0, err
	}
	p := path.Join(dir, u.Path)
	d := path.Dir(p)
	err = os.MkdirAll(d, os.ModePerm)
	f, err := os.Create(p)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err := io.Copy(f, r.Body)
	if err != nil {
		return 0, err
	}
	f.Sync()
	return int64(n), nil
}

func main() {
	dir := "."
	dataFile := "data"

	var wg sync.WaitGroup
	count := 10
	c := make(chan string)
	done := make(chan bool)

	for i := 1; i <= count; i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			Run(c, dir, done)
		}(i)
	}
	f, err := os.Open(dataFile)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		c <- s.Text()
	}
	close(done)
	wg.Wait()
}
