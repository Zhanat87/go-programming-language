// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		println("no args")
	} else {
		for _, arg := range files {
			absPath, _ := filepath.Abs("tasks/1/4/" + arg)
			f, err := os.Open(absPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg)
			f.Close()
		}
	}
}

func countLines(f *os.File, file string) {
	input := bufio.NewScanner(f)
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Printf("file with duplicated strings: %s\r\n", file)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
