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
	fmt.Println(newPerson("Jon", 27))
	fmt.Println(newPerson("Mikel", 24))
	fmt.Println(newPerson("John", 22))
	fmt.Println(&Person{name: "Sam", age: 22})

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
