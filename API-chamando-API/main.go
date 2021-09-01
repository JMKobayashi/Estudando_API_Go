package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://localhost:8000/")
	fmt.Println(res)
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
