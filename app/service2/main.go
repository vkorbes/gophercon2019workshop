package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serve)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Service 2 has received a request!")
	w.Write([]byte("Hello world!"))
}
