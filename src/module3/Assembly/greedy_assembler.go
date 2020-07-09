// Compass Quad 2020
package main

//GreedyAssembler takes a collection of strings and returns a genome whose
//k-mer composition is these strings. It makes the following assumptions.
//1. "Perfect coverage" -- every k-mer is detected
//2. No errors in reads
//3. Every read has equal length (k)
//4. DNA is single-stranded


// import "fmt"
// func main() {
// 	reads := []string{"AAT","ATG", "CAT", "CCA", "GAT", "GCC", "GGA", "GGG", "GTT", "TAA", "TGC", "TGG", "TGT"}
// 	// reads2 should have braces around it
// 	//prints TATAAATG
// 	//should print TAATGTA
// 	fmt.Println(GreedyAssembler(reads))
// 	//fmt.Println(GenerateRandomGenome(5))
// }



// NOTE: THE FOLLOWING CODE IS WRONG BECAUSE IT DOESN'T WORK WITH OVERLAPS :( 
func GreedyAssembler(reads []string) string {
	reads2 := DuplicateSlice(reads)
	k := len(reads2[0])
	genome := reads2[0]
	reads2 = reads2[1:]
	for _, read := range reads {
		if genome[len(genome)-(k-1):] == read[0:k-1] {
			genome = genome + string(read[k-1])
			reads2 = remove(read, reads2)
		}
	}

	for _, read := range reads2 {
		if genome[:k-1] == read[1:k] {
			genome = string(read[0]) + genome
			reads2 = remove(read, reads2)
		}
	}
	return genome
}

func DuplicateSlice(arr []string) []string {
	duplicate := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		duplicate[i] = arr[i]
	}
	return duplicate
}

func remove(item string, arr []string) []string {
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			index = i
			break
		}
	}
	myArr := append(arr[:index], arr[index+1:]...)
	return myArr
}
