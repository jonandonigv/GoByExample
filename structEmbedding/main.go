package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	c := container{
		base: base{num: 1},
		str:  "the number is 1",
	}

	fmt.Printf("co={num: %v, str: %v}", c.num, c.str)
	fmt.Println("also num:", c.base.num)
	fmt.Println("descrive:", c.describe())

	type describer interface {
		describe() string
	}

	var d describer = c
	fmt.Println("describer:", d.describe())
}
