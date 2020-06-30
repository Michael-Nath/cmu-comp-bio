package Jaccard

import "math"
// Write the function JaccardDistance() with the following input and output
// Input: two frequency maps sample1 and sample2
// Output: the Jaccard distance between sample1 and sample2
//thank you michael - anastasia
func JaccardDistance(sample1 map[string]int, sample2 map[string]int) float64{
    var Jaccard float64
    var minimum float64
    var maximum float64
    if len(sample1) > len(sample2){
        for val, _ := range sample1 {
            if sample1[val] != 0 && sample2[val] != 0 {
            minimum += math.Min(float64(sample1[val]), float64(sample2[val]))
          }
        }
        for val, _ := range sample1 {
            if sample1[val] != 0 && sample2[val] != 0 {
            maximum += math.Max(float64(sample1[val]), float64(sample2[val]))
          }
        }
    } else {
        for val, _ := range sample2 {
          if sample1[val] != 0 && sample2[val] != 0 {
          minimum += math.Min(float64(sample1[val]), float64(sample2[val]))
        }
        }
        for val, _ := range sample2 {
          if sample1[val] != 0 && sample2[val] != 0 {
            maximum += math.Max(float64(sample1[val]), float64(sample2[val]))
          }
        }
    }
    Jaccard = 1 - (minimum/maximum)
    return Jaccard
}
