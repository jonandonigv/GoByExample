package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	p := Person{name: name, age: age}
	return &p
}

func main() {
	fmt.Println(Person{"Bob", 20})
}
