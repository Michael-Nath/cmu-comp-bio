// Michael Nath
// Carnegie Mellon Computational Biology Summer Program
// Prologue Exercises

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Combination(1000, 2))
	fmt.Println(Permutation(1000, 2))
	fmt.Println(FactorialArray(10))
	fmt.Println(FibonacciArray(10))
	testArray := []int{56, 24, 48, 36, 20}
	fmt.Println(MinArray(testArray))
	fmt.Println(GCDArray(testArray))
	fmt.Println(IsPerfect(28))
	fmt.Println(NextTwinPrimes(31))
	// start := time.Now()
	// fmt.Println(NextPerfectNumber(8128))
	// elapsed := time.Since(start)
	// fmt.Println(elapsed)
}

//Factorial : calculates factorial value of a given int n
/*
Pseudocode:
Factorial(n)
	if n is negative -> throw an error saying it ain't gonna work
	if n is 0 or 1 -> return 1 (that's how factorials work)
	res <- 1
	for every integer i from n to 1:
		multiply res by i
	return value of res
*/
func Factorial(n int) int {
	if n < 0 {
		panic("Invalid Input for N")
	}
	if n < 1 {
		return 1
	}
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// Permutation : calculates permutation of k from n
// psuedocode:
// Permuation(n, k)
/* 	prod <- 1
for every integer i from n to n - k (descending order):
	prod is equal to itself multiplied by i
return prod
*/
func Permutation(n, k int) int {
	prod := 1
	for i := n; i > n-k; i-- {
		prod *= int(i)
	}
	if k > n {
		panic("k cannot be greater than n")
	}
	return prod
}

// Combination : calculates the combination of k from n
// pseudocode:
// Combination(n, k):
/*	prod <- 1
	smaller <- smaller of n-k and k
	larger <- greater of n-k and k
	for every integer p from n to the larger:
		prod becomes itself multiplied 
	return prod divided by factorial of the smaller number
*/



func Combination(n, k int) int {
	prod := 1
	smaller := int(math.Min(float64(n - k), float64(k)))
	larger := int(math.Max(float64(n - k), float64(k)))
	for i := n; i > larger; i-- {
		prod *= i
	}
	if k > n {
		panic("k cannot be greater than n")
	}
	return (prod / Factorial(smaller))
}

// FactorialArray : takes in an integer n and returns a slice of length n + 1 where the ith index is i!
// Psuedocode:
/*
	FactorialArray(n)
		factArray <- slice of size n + 1
		0th index of factArray <- 1
		1st index of factArray <- 1
		for every integer i from 2 to n:
			factArray of i <- factArray of (i - 1) * i
		return factArray
*/
func FactorialArray(n int) []int {
	factArray := make([]int, n+1)
	factArray[0] = 1
	factArray[1] = 1
	for i := 2; i < n+1; i++ {
		factArray[i] = i * factArray[i-1]
	}
	return factArray
}

// FibonacciArray : given an integer n, returns a slice of length n + 1 where the ith index represents the ith fibonacci number
// Pseudocode:
/*
	FibonacciArray(n):
		fibArray <- integer slice of length n + 1
		fibArray of 0 <- 1
		fibArray of 1 <- 1
		for every integer i from 2 to n:
			fibArray of i <- fibArray of (i - 1) + fibArray of (i - 2)
		return fibArray
*/
func FibonacciArray(n int) []int {
	fibArray := make([]int, n+1)
	fibArray[0] = 0
	fibArray[1] = 1
	for i := 2; i < n+1; i++ {
		fibArray[i] = fibArray[i-1] + fibArray[i-2]
	}
	return fibArray
}

// MinArray : given an integer array, returns the minimum value
// Pre-Condition: the array isn't empty
// Pseudocode:
/*
	MinArray(n):
		currentMinimum <- the 0th element of n
		for every integer i from 1 to the length of the array - 1:
			if n of i is smaller than currentMinimum:
				currentMinimum <- n of i
		return currentMinimum
*/
func MinArray(n []int) int {
	if len(n) == 0 {
		panic("pass in a non-empty array")
	}
	currentMinimum := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] < currentMinimum {
			currentMinimum = n[i]
		}
	}
	return currentMinimum
}

// GCDArray : given an array of integers, returns the greatest common divisors between all the integers
// Pseudocode:
/* GCDArray(n):
smallestValue <- value of MinArray on n
for every factor p of smallestValue in descending order:
	if p is 1 -> return 1
	for every index i in n:
		if p is not a factor of n of i:
			move onto the next p
	if p is a factor of all integers in n:
		return p
*/
func GCDArray(n []int) int {
	smallestValue := MinArray(n)
	answer := 1
	for p := smallestValue; p > 0; p-- {
		isFactor := true
		for i := 0; i < len(n); i++ {
			if n[i]%p != 0 {
				isFactor = false
				break
			}
		}
		if isFactor {
			answer = p
			break
		}
	}
	return answer
}

// IsPerfect : takes in an integer n and returns a true/false value depending on whether n is a perfect number
// pseudocode:
/*
	IsPerfect(n):
		sumPropDivisors <- 0
		for every proper divisor p of n:
			sumPropDivisors <- sumPropDivisor + p
		if sumPropDivisor equals n:
			return true
		otherwise return false
*/
func IsPerfect(n int) bool {
	sumPropDivisors := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sumPropDivisors += i
			sumPropDivisors += (n / i)
		}
	}
	return sumPropDivisors == n
}

// NextPerfectNumber : given an integer n, returns the smallest perfect number greater than n
// psuedocode:
/*
NextPerfectNumber(n):
	p <- n + 1
	while p is not a perfect number:
		p <- p + 1

	return p
*/
func NextPerfectNumber(n int) int {
	p := n + 1
	for !IsPerfect(p) {
		p++
	}
	return p
}

// IsPrime : Given an integer p, returns whether or not the integer is a prime number
func IsPrime(p int) bool {
	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		if p%k == 0 {
			return false
		}
	}
	return true
}

// NextTwinPrimes : given an integer n, returns a pair of twin primes that are the least possible pair greater than n
// pseudocode:
func NextTwinPrimes(n int) (int, int) {
	starter := n + 1
	for !IsPrime(starter) || !IsPrime(starter+2) {
		starter++
	}
	return starter, starter + 2
}
