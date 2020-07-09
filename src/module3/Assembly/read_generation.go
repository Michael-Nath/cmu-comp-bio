package main

import (
	"math/rand"
)

//SimulateReadsClean takes a genome along with a read length and a probability.
//It returns a collection of strings resulting from simulating clean reads,
//where a given position is sampled with the given probability.
func SimulateReadsClean(genome string, readLength int, probability float64) []string {
	cleanReads := make([]string, 0)
	for i := 0; i <= (len(genome) - readLength); i++ {
			if rand.Float64() < probability {
				cleanReads = append(cleanReads, genome[i:i+readLength])
			}
	}
	return cleanReads
}

// func SimulateReadsMessy(genome string, readLength int, probability, substitutionErroRate, insertionErrorRate, deletionErrorRate float64) []string{
// 		reads := make([]string, 0)
// 		if substitutionErrorRate + insertionErrorRate + deletionErrorRate >= 1 {
// 				panic("Error: Sean is a stickler :)")
// 		}
// 		for i := 0; i < len(genome) - readLength + 1; i++ {
// 				x := rand.Float64()
// 				if x < probability {
// 						currString := ""
// 						currPosition := i
// 						for len(currString) < readLength && currPosition < len(genome){
// 								y := rand.Float64()
// 								if y < substitutionErrorRate {
// 										sym := RandomDnaSymbol()
// 										for sym == genome[currPosition] {
// 												sym = RandomDnaSymbol()
// 										}
// 										currString += string(sym)
// 										currPosition++
// 								} else if y < substitutionErrorRate + insertionErrorRate {
// 										currString += string(RandomDnaSymbol())
// 								} else if y < substitutionErrorRate + insertionErrorRate + deletionErrorRate {
// 										currPosition++
// 								} else {
// 										currString += string(genome[currPosition])
// 										currPosition++
// 								}
// 						}
// 						if len(currString) == readLength {
// 								reads = append(reads, currString)
// 						}
// 				}
// 		}
// 		return reads
// }

//NOTE: we will simulate messy reads together.

//SimulateReadsMessy takes a genome along with a read length, a probability,
//and error rates for substitutions, insertions, and deletions.
//It returns a collection of reads resulting from simulating clean reads,
//where a given position is sampled with the given probability, and
//errors occur at the respective rates.
