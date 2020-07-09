package main

import "math/rand"
// import "time"

//GenerateRandomGenome takes a parameter genomeLength and returns
//a random DNA string of this length where each nucleotide has equal probability.

// Slower Version
// func GenerateRandomGenome(genomeLength int) string {
// 	var genome string
// 	//FILL IN YOUR CODE HERE
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < genomeLength; i++ {
// 		n := rand.Intn(4) + 1
// 		if n == 1 {
// 			genome += "A"
// 		} else if n == 2 {
// 			genome += "C"
// 		} else if n == 3 {
// 			genome += "G"
// 		} else if n == 4 {
// 			genome += "T"
// 		}
// 	}
//
// 	return genome
// }

func GenerateRandomGenome(genomeLength int) string {
	//make an array of symbols
	symbols := make([]byte, genomeLength)

	for i := 0; i < genomeLength; i ++{
		symbols[i] = RandomDnaSymbol()
	}
	return string(symbols)
}

func RandomDnaSymbol() byte{
	i := rand.Intn(4)
	symbols := []byte{'A', 'T', 'G', 'C'}
	return symbols[i]
}
