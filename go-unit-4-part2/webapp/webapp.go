package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
)

func handler(w http.ResponseWriter, r *http.Request) {

	var i = false
	var j = false
	var k = false

	searchkeyword := r.URL.Query().Get("word")

	fruits := []string{"apple", "banana", "chikoo", "grapes", "pear", "mango", "kiwi", "watermelon"}
	veggies := []string{"beans", "carrot", "drumstick", "brinjal", "tomato", "beetroot"}
	pulses := []string{"rice", "wheat", "ragi", "jowar"}

	i = slices.Contains(fruits, searchkeyword)

	j = slices.Contains(veggies, searchkeyword)

	k = slices.Contains(pulses, searchkeyword)

	if i == true {
		fmt.Fprintf(w, searchkeyword)
		fmt.Fprintf(w, " is a fruit\n")
	}

	if j == true {
		fmt.Fprintf(w, searchkeyword)
		fmt.Fprintf(w, " is a vegetable\n")
	}

	if k == true {
		fmt.Fprintf(w, searchkeyword)
		fmt.Fprintf(w, " is a pulse\n")
	}

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//http://localhost:8080/?word=mango
