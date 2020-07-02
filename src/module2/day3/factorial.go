package main
import "fmt"

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(Fib(60))
}

func Fib(n int) int {
	if n < 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
