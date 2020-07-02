package Alignment
//LongestSharedSubstring takes two strings as input.  It returns the shared substring
//of the input strings of maximum length, via a brute-force algorithm.
//If multiple answers exist, it returns only one.


// start with the smaller string
// create a current substring variable to hold the current largest substring
// go through the smaller and larger string together, and then take continious chunks of it
// compare the chunks and see if they're equal. If they are, then


func LongestSharedSubstring (s1, s2 string) string {
    smallerString := ""
    longerString := ""
    if len(s1) < len(s2) {
      smallerString = s1
      longerString = s2
    } else {
      smallerString = s2
      longerString = s1
    }
    possibilities := make([]string, 0)
    // var currentSubstring string[i:j]
    for i := 0; i < len(smallerString); i++ {
      for j := i + 1; j <= len(smallerString); j++ {
        possibilities = append(possibilities, smallerString[i:j])
      }
    }
    counter := 0;
    answer := ""
    for i := 0; i < len(longerString); i++ {
      for j := i + 1; j <= len(longerString); j++ {
        for _, sub := range possibilities {
          if longerString[i:j] == sub && len(sub) > counter{
            answer = sub;
            counter = len(sub)
          }
        }
      }
    }
    return answer
}


/*
if len(s1) > len(s2) {
  temp := s1
  s1 = s2
  s2 = temp
}
for L := len(s1); L >= 1; L-- { // starting from longest possible substr length
  for i := 0; i+L <= len(s1); i++ {
    str := s1[i : i+L]
    for j := 0; j+L <= len(s2); j++ {
      if str == s2[j:j+L] {
        return str
      }
    }
  }
}
return ""
*/
