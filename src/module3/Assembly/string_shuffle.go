// Compass Quad 2020

package main
import "math/rand"
import "time"
//ShuffleStrings takes a collection of strings patterns as input.
//It returns a random shuffle of the strings.
func ShuffleStrings(patterns []string) []string {
	rand.Seed(time.Now().UnixNano())
	// creates a random permutation
	myPerms := rand.Perm(len(patterns))
	shuffledStrings := make([]string, len(patterns))
	// assigns a string from patterns to a random spot in the shuffledStrings array (determined by the values in myPerms)
	for i := 0; i < len(myPerms); i++ {
		shuffledStrings[myPerms[i]] = patterns[i]
	}
	return shuffledStrings
}
