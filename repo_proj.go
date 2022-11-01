package main

import "fmt"

type Person struct {
	name string
	Age  string
}

func main() {
	fmt.Println(Person{
		"Patrice", "33",
	})

}
