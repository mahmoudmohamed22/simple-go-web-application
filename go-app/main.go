package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	fmt.Println("Success is the sum of small efforts, repeated day-in and day-out. :) ")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello , I am Mahmoud Abdelwahab Elghonemy\n")
		fmt.Fprintf(w,"\n\nSuccess is the sum of small efforts, repeated day-in and day-out. :) ")

	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
