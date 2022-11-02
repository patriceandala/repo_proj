package main

import "fmt"

type Person struct {
	name string
	Age  string
}

func main() {
	fmt.Println(Person{
		"Test", "Reb",
	})
	fmt.Println(Person{
		"Other", "Ish",
	})

}
