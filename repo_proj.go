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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := fmt.Fprintf(w, "What's good")
		if err != nil {
			return
		}
	})

}
