package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// practiceLoop()
	// roulette()
	castToFloat()
}

// for文の練習
func practiceLoop() {
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

// おみくじ
func roulette() {
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		rand := rand.Intn(6) + 1
		if rand == 1 {
			fmt.Println("凶")
		} else if rand == 2 || rand == 3 {
			fmt.Println("吉")
		} else if rand == 4 || rand == 5 {
			fmt.Println("中吉")
		} else if rand == 6 {
			fmt.Println("大吉")
		}
	}
}

func castToFloat() {
	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	flo := float64(avg)
	if flo > 4.5 {
		fmt.Printf("%v", flo)
		println("good")
	} else {
		fmt.Printf("%v", flo)
		println("bad")
	}
}
