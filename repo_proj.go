package main

import (
	"fmt"
	"net/http"
)

type Person struct {
	name string
	Age  string
}

func main() {
	fmt.Println(Person{
		"Test", "Reb",
	})
	fmt.Println(Person{
		"Other", "IrIsh",
	})
	fmt.Println(Person{
		"Other", "Ish",
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := fmt.Fprintf(w, "What's good")
		if err != nil {
			return
		}
	})

}
