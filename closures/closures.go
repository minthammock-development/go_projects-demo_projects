package main
import "fmt"
func intSec() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	netInt := intSec()

	fmt.Println(netInt())
	fmt.Println(netInt())
	fmt.Println(netInt())
	fmt.Println(netInt())

	newInts := intSec()
	fmt.Println(newInts())
}