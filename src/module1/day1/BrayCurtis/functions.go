package BrayCurtis
import "math"
// Write the function BrayCurtisDistance() with the following input and output
// Input: two frequency maps sample1 and sample2
// Output: the Bray-Curtis distance between sample1 and sample2

func BrayCurtisDistance(sample1 map[string]int, sample2 map[string]int) float64 {
    var BrayCurtis float64
    avgTotal := float64(Total(sample1) + Total(sample2)) / 2.0
    var minimum float64
    for val, _ := range sample1 {
        minimum += math.Min(float64(sample1[val]), float64(sample2[val]))
    }
    BrayCurtis = 1 - (minimum/avgTotal)
    return BrayCurtis
}

func Total(sample1 map[string]int) int {
  total := 0
  for _, val := range sample1 {
    total += val
  }
  return total;
}
