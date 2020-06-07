package main
import "fmt"

func main() {
	// // fmt.Println(Factorial(5))
	fmt.Println(Combination(1000, 2))
	fmt.Println(Permutation(1000,2))
	fmt.Println(Combination(1000, 998))
	// fmt.Println((Factorial(1000 - 998) * Factorial(998)))
	// fmt.Println(Factorial(998))
}

//Factorial : calculates factorial value of a given int n
func Factorial(n int) float64 {
	if n < 0 {panic("Invalid Input for N")}
	if n < 1 {return 1}
	res := 1.0 
	for i := 1; i <= n; i++ {
		res *= float64(i)
	}
	return res
}

// Permutation : calculates permutation of k from n
func Permutation(n, k int) float64 {
	if k > n {panic("k cannot be greater than n")}
	return (Factorial(n) / Factorial(n - k))
}

// Combination : calculates the combination of k from n
func Combination(n, k int) float64 {
	if k > n {panic("k cannot be greater than n")}
	return Factorial(n) / ((Factorial(n - k) * Factorial(k)))
}