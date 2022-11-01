package main

import "fmt"

type Person struct {
	name string
	Age  string
}

func (p Person) printName() {
	fmt.Println(p.name)
}

func main() {
	p := Person{
		name: "Patrice",
		Age:  "32",
	}
	p.printName()
}
