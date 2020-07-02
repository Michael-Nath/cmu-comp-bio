package Alignment
import (
  "math"
  )
//Change takes an amount of money along with a collection of denominations as a []int.
//It returns the minimum number of coins needed to change the money given the denominations.

func Change(n int, coins []int) int {
    mini := make([]int, n+1) //minimum # of coins to make i
    if n == 0 {return 0}
    mini[0] = 0
    for _,coin := range coins {
      if n == coin {return 1}
      mini[coin] = 1
    }
    globalCounter := math.MaxInt32
    for i := 1; i <= n; i++ {
      counter := math.MaxInt32
      for j := 0; j < len(coins); j++ {
        if (i - coins[j]>=0 && 1 + mini[i - coins[j]] < counter) {counter = 1 + mini[i - coins[j]]}
      }
      globalCounter = counter
      mini[i] = counter
    }
    if (globalCounter == math.MaxInt32) {return (n / coins[0]) + 1}
    return mini[n]
}

// func Min(nums []int) int {
//   smallest := nums[0]
//   for i := 1; i < len(nums); i++ {
//     if nums[i] < smallest {smallest = nums[i]}
//   }
//   return smallest
//  }
