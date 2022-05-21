package main
import (
	"fmt"
)
type person struct {
	name string
	age int 
}

func newPersonPointer(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
func main(){
	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "Alice", age: 40})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 50})
	fmt.Println(newPersonPointer("Jon"))

	s := person{name: "Sean", age: 76}
	s.name = "Me"
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println((sp.age))
}