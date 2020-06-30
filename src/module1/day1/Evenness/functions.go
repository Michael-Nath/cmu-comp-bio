package Evenness

// Write a function SimpsonsIndex() with the following input and output
// Input: a frequency map sample
// Output: the Simpson’s index of sample. That is, the decimal sum of (n/N) squared
// over all species where n = # of individuals found for a given string/species
// and N is the total # of individuals.
func SimpsonsIndex (sample map[string]int) float64{
    var output float64 = 0
    var Total int = 0 //Total is big N
    for _, val := range sample {
      output += float64(val) * float64(val)
      Total += val
    }
    return output / (float64(Total) * float64(Total))
}
