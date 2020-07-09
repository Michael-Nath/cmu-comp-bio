//Compass Quad 2020
package main
import "fmt"
func main() {
  fmt.Println("First, greedy assembly.")
  RunGreedyAssembly()
}

func RunGreedyAssembly() {
  genomeLength := 100
  k := 10
  // first generate a random genome
  originalGenome := GenerateRandomGenome(genomeLength)

  // form it's k-mer composition
  reads := KmerComposition(originalGenome, k)

  // now call our greedy assembler
  assembledGenome := GreedyAssembler(reads)

  kmers := KmerComposition(assembledGenome, k)

  if (StrSliceEq(kmers,reads)) {
    fmt.Println("Mission Accomplished!")
  }

}
