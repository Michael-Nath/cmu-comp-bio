// Compass Quad
// Day 1 FrequencyMap Exercise


package FrequencyMap

// Write function FrequencyMap() with the following input and output
// Input: one collection of strings patterns
// Output: a frequency map of strings to their # of counts in patterns

func FrequencyMap(patterns []string) map[string]int {
  // initialize a frequency map
  // go through the patterns array
  // add a given pattern if it's not there in the map
  // otherwise add 1 to it.
  freq := make(map[string]int)
  for j := 0; j < len(patterns); j++ {
  // for _,val := range patterns (this is what Marcus suggested) {
      num := patterns[j]
      freq[num]++
  }
  return freq
}

/* josh's solution:
package FrequencyMap

// Write function FrequencyMap() with the following input and output
// Input: one collection of strings patterns
// Output: a frequency map of strings to their # of counts in patterns

func FrequencyMap(patterns []string) map[string]int{
  freqMap:= make(map[string] int)
  for _,val:= range patterns{
    freqMap[val]++
  }
  return freqMap
}
*/
