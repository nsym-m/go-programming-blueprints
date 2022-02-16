package main

import (
	"fmt"
	"strconv"
)

func main() {
	// for
	for i := 1; i <= 100; i++ {
		// if
		if i%2 == 0 {
			// cast
			fmt.Printf("%s-偶数\n", strconv.Itoa(i))
		} else {
			fmt.Printf("%s-奇数\n", strconv.Itoa(i))
		}
		// switch
		switch i % 2 {
		case 0:
			fmt.Printf("%s-偶数\n", strconv.Itoa(i))
		default:
			fmt.Printf("%s-奇数\n", strconv.Itoa(i))
		}
	}
}
