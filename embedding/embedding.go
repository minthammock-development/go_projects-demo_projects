package main
import  "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("Base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{num: 1},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also numn: ", co.base.num)
	fmt.Println("describe method: ", co.describe())
	fmt.Println("long way describe method: ", co.base.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("desctiber: ", d.describe())
}