package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	http.HandleFunc("/", serve)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Service 2 has received a request!")
	message, err := ioutil.ReadFile("message.txt")
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(message)
}
