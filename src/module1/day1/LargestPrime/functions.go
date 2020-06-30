package LargestPrime

import (
    "math"
)
//Write a function to find a largest prime number.
//LargestPrime takes an integer n as input.
//It panics if n < 2. Otherwise, it returns the largest prime
//factor of n.

func LargestPrimeFactor(n int) int {
  // for loop from 1 to n
  // find the number factors
  // check if that number is also a prime number
  // let's start out with 2 since it's the first prime number
  var current int
  if IsPrime(n) {
      return n
  } else {
    if (n == 2) {current = 2}
    if (n == 1) {current = 1}
    if (n <= 0) {panic("Error: input to LargestPrime is not positive")}
    for num := n; num > 1; num-- {
      if n % num == 0 && IsPrime(num){
        current = num;
        return current
      }
    }
  }
  return current
}

func IsPrime(p int) bool {
  for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
    if p % k == 0 {
      return false
    }
  }
  //if we survive then p is primeArray
  return true
}
