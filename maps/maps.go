package main
import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 14
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1 is this: ",v1)
	fmt.Println("len of m: ", len(m))

	delete(m, "k2")
	fmt.Println("m is now: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int {"foo": 1, "bar": 2}
	fmt.Println("map n looks like this: ", n)
}