// Compass Quad 2020 Modified Jaccard



package Jaccard

import (
  "math"
)
// Write the function JaccardDistance() with the following input and output
// Input: two frequency maps sample1 and sample2
// Output: the Jaccard distance between sample1 and sample2
//thank you michael - anastasia
func JaccardDistance(sample1 map[string]int, sample2 map[string]int) float64{
    var Jaccard float64
    var minimum float64
    var maximum float64
    seen := make(map[string]int, 0)
    // first we create a map of dna we've seen so far
    // we go through the first sample, and if we haven't seen the dna before, map[dna] = 1
    // we go through the second sample, and we check if map[dna] = 1
        for val, _ := range sample1 {
            minimum += math.Min(float64(sample1[val]), float64(sample2[val]))
            maximum += math.Max(float64(sample1[val]), float64(sample2[val]))
            seen[val] = 1
        }
        for val, _ := range sample2 {
            if seen[val] != 1 {
                minimum += math.Min(float64(sample1[val]), float64(sample2[val]))
                maximum += math.Max(float64(sample1[val]), float64(sample2[val]))
            }
        }
    Jaccard = 1 - (minimum/maximum)
    return Jaccard
}
