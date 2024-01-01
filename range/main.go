package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	user := map[string]string{"name": "jon", "lastname": "galdos"}
	for _, c := range user {
		fmt.Println(c)
	}

	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	for k, v := range kvs {
		fmt.Println("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
