// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"net/url"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, link := range os.Args[1:] {
		go fetch(link, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(link string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(link)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", link, err)
		return
	}
	u, err := url.Parse(link)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", link, err)
		return
	}
	err = ioutil.WriteFile("./tasks/1/10/" + u.Host + ".txt", body, 0666)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", link, err)
		return
	}
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", link, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%s", secs, link)
}

//!-
