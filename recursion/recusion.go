package main
import "fmt"

func fact(n int) int {
	if n == 0 {
			return 1
	}
	return n* fact(n-1)
}

func main() {
	fmt.Println(fact(7))
	var fib func(n int) int 
	
	fib = func(n int) int {
		if n == 0 {
			return 1
			}
		return fact(n-1) + n
	}
	fmt.Println(fib(7))
}

