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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := fmt.Fprintf(w, "What's good, n word -Stinkmener. Now reap what you sow nigga")
		if err != nil {
			return
		}
	})

}
