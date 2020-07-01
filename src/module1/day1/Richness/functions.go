package Richness

// Write function Richness() with the following input and output
// Input: a frequency map sample
// Output: the richness of sample. (i.e., the number of keys in the map
// corresponding to nonzero values.)

//Richness takes a frequency map. It returns the richness of the frequency map
//(i.e., the number of keys in the map corresponding to nonzero values.)


//From sample
func Richness(sample map[string]int) int {
	c := 0
	for _, val := range sample {
		if val > 0 {
			c++
		}
	}
	return c
}
