// Compass Quad 2020
package main
//KmerComposition returns the k-mer composition (all k-mer substrings) of a given genome.
func KmerComposition(genome string, k int) []string {
		AllKmers := make([]string, 0)
		for i := 0; i < len(genome)-(k-1); i++ {
				var str string
				for j := i; j < i + k; j++ {
						str += string(genome[j])
				}
				AllKmers = append(AllKmers, str)
		}
		return AllKmers
}

// AATTC
