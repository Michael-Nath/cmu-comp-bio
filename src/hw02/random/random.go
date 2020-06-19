// Michael Nath
// Carnegie Mellon Computational Biology Summer Program
// Chapter 2 Exercises

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(RelativelyPrimeProbability(2, 4, 10))
	fmt.Println(BirthdayParadox(50, 100000))
	test := []int{1473, 2856, 9830, 1789, 4468, 9830}
	fmt.Println(ComputePeriodLength(test))
	fmt.Println(CountNumDigits(123456789))
	fmt.Println(SquareMiddle(8100, 4))
	fmt.Println(GenerateMiddleSquareSequence(70, 2))
}

// =====================================
// Monte Carlo Relative Prime
// Write a function RelativelyPrimeProbability that takes three int as input(x,
// y, numPairs). Uniformly generate pairs of integers between x and y, inclusive
// (numPairs total), and return an estimate of probability of the pairs being
// relatively prime(type float64).

// Your program needs to panic if any of the input is non-positive.

// numPairs <= 100000. For all inputs, either 1 <= x, y <= 1000, or
// x = 1 and y <= 1000000.

// The PRNG will be seeded by the autograder and you are
// free to use it in whatever way you like. We will grade your implementation
// by checking that your output is within a certain amount of the correct answer.
// If your implementation is correct, it has at least 99.9999%
// probability to pass. If you believe your implementation
// is correct but autograder have it wrong, buy a lottery today as it's your
// (super) lucky day!

// Sample Input: (2, 4, 10)
// Sample Output: 0.5
// Note: You can get anything from 0 to 1. The actual probability for these x and y is 4/9 = .444.
// =====================================

// EuclidGCD : given two integers, returns their GCD
func EuclidGCD(x, y int) int {
	if x < 0 || y < 0 {
		panic("Invalid arguments. Must be non-negative")
	}
	if x == 0 {
		return y
	}

	if y == 0 {
		return x
	}

	return EuclidGCD(y, x%y)
}

// RelativelyPrimeProbability : takes three int as input(x,
// y, numPairs). Uniformly generate pairs of integers between x and y, inclusive
// (numPairs total), and return an estimate of probability of the pairs being
// relatively prime(type float64).
func RelativelyPrimeProbability(x, y, numPairs int) float64 {
	count := 0
	for i := 0; i < numPairs; i++ {
		rand.Seed(time.Now().UnixNano())
		firstInt := rand.Intn(y-x+1) + x
		secondInt := rand.Intn(y-x+1) + x
		if EuclidGCD(firstInt, secondInt) == 1 {
			count++
		}
	}
	return float64(count) / float64(numPairs)
}

// =====================================
// Repeat Sequences
// Write a function HasRepeat that takes a slice of integers as input and
// returns a boolean indicating whether there is a repeated value (a value that
// appears twice or more in the slice).

// Length of slice <= 500. Elements in slice <= 100000000 in absolute value.
// Sample Input: [1, 2, 3, 4, 5, 6, 7, 1, 2, 3]
// Sample Output: true
// =====================================

// HasRepeat : takes a slice of ints and returns bool indicating whether there's a repeated value.
func HasRepeat(elems []int) bool {
	for i := 0; i < len(elems); i++ {
		for j := i + 1; j < len(elems); j++ {
			if elems[j] == elems[i] {
				return true
			}
		}
	}
	return false
}

// =====================================
// Monte Carlo Birthday Paradox
// Write a function BirthdayParadox that takes two integers as input(numPeople, numTrials).
// Simulate numPeople random birthdays and estimate the probability that at least two
// of them are the same (sample numTrials times total). Return the estimate
// as a float64.

// Your program should panic if either of the inputs is non-positive.

// numPeople <= 50, numTrials <= 100000. We have the same guarantee regarding
// testing your answer for correctness as in the Monte Carlo relatively prime problem.

// Sample Input: numPeople = 2, numTrials = 1000
// Sample Output: 0.003
// Note: You can get anything from 0 to ~0.01. The actual probability is
// 1/365 = 0.0027...

// =====================================

// BirthdayParadox : refer to above
func BirthdayParadox(numPeople, numTrials int) float64 {
	if numPeople <= 0 || numTrials <= 0 {
		panic("Invalid arguments; must be positve")
	}
	count := 0.0
	for i := 0; i < numTrials; i++ {
		peepArray := make([]int, numPeople)
		for peep := 0; peep < numPeople; peep++ {
			rand.Seed(time.Now().UnixNano())
			day := rand.Intn(365) + 1
			peepArray[peep] = day
		}
		if HasRepeat(peepArray) {
			count++
		}
	}
	return count / float64(numTrials)
}

// =====================================
// Period Length
// Write a function ComputePeriodLength that takes a slice of int as input and
// check if you can decide the period of the input sequence. If no, return 0;
// If yes, return the period.

// Length of slice <= 500. Elements in slice <= 100000000 in absolute value.
// Sample Input: [1, 2, 3, 4, 5, 6, 7, 1, 2, 3]
// Sample Output: 7
// =====================================

// ComputePeriodLength : takes a slice of int as input and returns period
func ComputePeriodLength(elems []int) int {
	for i := 0; i < len(elems); i++ {
		temp := 0
		for j := i + 1; j < len(elems); j++ {
			temp++
			if elems[j] == elems[i] {
				return temp
			}
		}
	}
	return 0
}

//=====================================
// Number of Digits
// Write a function CountNumDigits that takes an integer x(type int) as input,
// and return number of digits in x. By convention, 0 have 1 digit, and (-x)
// have same number of digits as x, for any positive integer x.

// You should not rely on any functions from imported packages for this problem.
// Input number <= 10^10.
// Sample Input: 12345
// Sample Output: 5
//=====================================

// CountNumDigits : given an integer x, returns the number of digits in x
func CountNumDigits(x int) int {
	count := 1
	for x/10 != 0 {
		count++
		x /= 10
	}
	return count
}

// =====================================
// SquareMiddle
// Write a function SquareMiddle that takes two ints: x and numDigits, and
// return the number formed by middle numDigits of x squared.
// Additionally, we ask you to perform sanity checks. If numDigits is not even,
// or if x has more than 2*numDigits of digit, or either input is negative,
// your function should panic.

// For all valid data: 0 <= x < 10^8, 0 < numDigits <= 8
// Your function need to panic if the either of the input falls outside
// of valid range, OR if x is no less than 10^numDigits, OR numDigits is not even.
// Sample Input 1: x = 55, numDigits = 2
// Sample Output 1: 2 (55 * 55 = 3025)
// Sample Input 2: x = 55, numDigits = 3
// Sample Output 2: None(You should panic!)
// =====================================

// SquareMiddle : refer to above
func SquareMiddle(x, numDigits int) int {
	stringX := strconv.Itoa(x * x)
	if numDigits%2 != 0 || CountNumDigits(x) > 2*numDigits || x < 0 || numDigits <= 0 {
		panic("Error! Please check your arguments")
	}
	skipper := numDigits / 2
	beginner := skipper
	end := len(stringX) - skipper
	fmt.Println(stringX)
	if len(stringX) == (2*numDigits)-1 {
		beginner = skipper - 1
	}
	fmt.Println(len(stringX))
	res, err := strconv.Atoi(stringX[beginner:end])
	if err != nil {
		panic(err)
	}
	return res
}

// =====================================
// SquareMiddle Generator
// Now that you have every tool you need, let's implement the actual generator.
// Write a function GenerateMiddleSquareSequence that takes two ints: seed and
// numDigits, and output the generated sequence right after first repeated
// element appears(stop at this point; This is when you know the period of your
// sequence).

// 0 <= seed < 10^numDigits. 2 <= numDigits <= 6. We guarantee the output
// sequence has less than 500 elements.
// Sample Input: seed = 70, numDigits = 2
// Sample Output: [70, 90, 10, 10]
// =====================================

// GenerateMiddleSquareSequence : refer to above
func GenerateMiddleSquareSequence(seed, numDigits int) []int {
	arr := make([]int, 0)
	num := seed
	arr = append(arr, seed)
	for ComputePeriodLength(arr) == 0 {
		num = SquareMiddle(num, numDigits)
		arr = append(arr, num)
	}
	return arr
}
