package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
func main() {
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxx")
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Println(err)
	}
}
