package MultipleEvenness

import (
	"Metagenomics/Evenness"
)

// Write the function SimpsonsMatrix() with the following input and output
// Input: an array of frequency maps samples
// Output: an slice  i-th element is the Simpson's index of the i-th frequency map

func SimpsonsMatrix(AllMaps []map[string]int) []float64 {
		SimpsonSlice := make([]float64,len(AllMaps))
		for i := 0; i < len(AllMaps); i++ {
			SimpsonSlice[i] = Evenness.SimpsonsIndex(AllMaps[i])
		}
		return SimpsonSlice
}

func SimpsonsIndex (sample map[string]int) float64{
    var output float64 = 0
    var Total int = 0 //Total is big N
    for _, val := range sample {
      output += float64(val) * float64(val)
      Total += val
    }
    return output / (float64(Total) * float64(Total))
}
