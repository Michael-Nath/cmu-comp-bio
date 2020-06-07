package main
import "fmt"

func main() {
	fmt.Println(calcSumSquares(100))
}
// sums the squares of ints from 1 to n inclusively
func calcSumSquares(n int) int {
	// inits a variable sum that keeps track of sum of ints
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (i * i)
	}

	return sum
}