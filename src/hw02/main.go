package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(LoadedDie())
	fmt.Println(EuclidGCD(24,31))
}

// LoadedDie : simulates a weighted die where the probability of getting 3 is 0.5
func LoadedDie() int {
	result := rand.Intn(10) + 1
	choices := []int{1, 2, 4, 5, 6}
	if result <= 5 {
		return 3
	}
	return choices[result-6]

}
// EuclidGCD : given two integers, returns their GCD
func EuclidGCD(x,y int) int {
	if x == 0 {
		return y
	}

	if y == 0 {
		return x
	}

	return EuclidGCD(y, x % y);
}