package main
import "fmt"
func main() {
	var a [5]int
	fmt.Println("empty: ", a)

	a[4] = 100
	fmt.Println("set value: ", a)
	fmt.Println("get value: ", a[4])
	fmt.Println("len a: ",len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("declared: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = j
		}
	}
	fmt.Println("twoD looks like this: ", twoD)  
}