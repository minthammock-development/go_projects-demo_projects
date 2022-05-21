package main
import "fmt"
func zeroval(ival int) int{
	ival = 0
	return ival
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i )
	fmt.Println("initial pointer: ", &i)

	zeroval(i)
	fmt.Println("after zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zero pointer: ", i)

	fmt.Println("pointer itself: ", &i)

	i = 5
	fmt.Println(i)
	fmt.Println(&i)

}