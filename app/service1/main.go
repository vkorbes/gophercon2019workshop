package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", serve)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Service #1 has received a request!")
	message := []byte(fmt.Sprint("Service #2 says: ", howdy()))
	w.Write(message)
}

func howdy() string {
	resp, err := http.Get("http://dep-svc2:8080/")
	// resp, err := http.Get("http://localhost:8081/")
	if err != nil {
		log.Fatalln(err)
	}
	message, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(message)
}