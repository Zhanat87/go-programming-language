// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for i := 0; i < len(os.Args); i++ {
		if i == 0 {
			s += fmt.Sprintf("command name: %s\r\n", os.Args[i])
		} else {
			s += fmt.Sprintf("#%d. %s\r\n", i, os.Args[i])
		}
	}
	fmt.Println(s)
}

//!-
