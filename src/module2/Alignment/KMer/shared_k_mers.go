package Alignment

//CountSharedKmers takes two strings and an integer k. It returns the number of
//k-mers that are shared by the two strings.
// use the FrequencyMap function to create occurneces for each string.
// go through the shorter map.
// see if a given value is in the larger map.
// if it is, then compute the minimum and add it to a variable that stores the sum of all minima.
func CountSharedKmers(s1, s2 string, k int) int {
    s1Map := FrequencyMap(s1, k)
    s2Map := FrequencyMap(s2, k)
    shorterMap := make(map[string]int)
    longerMap := make(map[string]int)
    if len(s1Map) < len(s2Map) {
      shorterMap = s1Map
      longerMap = s2Map
    } else {
      shorterMap = s2Map
      longerMap = s1Map
    }
    answer := 0
    for key, _ := range shorterMap {
      if longerMap[key] != 0 {
        answer += Min(shorterMap[key], longerMap[key])
      }
    }
    return answer
}

func Min(x, y int) int {
  if (x < y) {
    return x
  }
  return y
}

func FrequencyMap(text string, k int) map[string]int {
	// map declaration is analogous to slices
	// (we don't need to give an initial length)
	freq := make(map[string]int)
	n := len(text)
	for i := 0; i < n-k+1; i++ {
		pattern := text[i : i+k]
		// if freqMap[pattern] doesn't exist, create it.  How do we do this?
		/*
		   // approach #1
		   _, exists := freq[pattern]
		   if exists == false {
		     // create this value
		     freqMap[pattern] = 1
		   } else {
		     // we already have a value in the map
		     freqMap[pattern]++
		   }
		*/
		// approach #2
		// this says, if freqMap[pattern] exists, add one to it
		// if freqMap[pattern] doesn't exist, create it with a default value of 0, and add 1.
		freq[pattern]++
	}
	return freq
}










// AT GTTATA
// ATCGTCC

// 4 matches
