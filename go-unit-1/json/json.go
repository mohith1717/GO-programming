package main

import (
	"encoding/json"
	"fmt"
)

func main() {
 
     type Movie struct {
                         Title string
                         Year int `json:"released"`
                         Color bool `json:"color,omitempty"`
                         Actors []string
                       }

     var movies = []Movie{
                           {Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
                           {Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
                           {Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
                         }

     data, err := json.Marshal(movies)

     if err != nil {
                     fmt.Printf("JSON marshaling failed: %s", err)
                   }

     fmt.Printf("%s\n\n", data)
}
