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
	fmt.Println(c)
	fmt.Println(base{num: 1}.describe())
	fmt.Println(c.describe())
}
