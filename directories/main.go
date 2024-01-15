package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")

}
