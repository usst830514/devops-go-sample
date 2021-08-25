package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello,"+req.URL.Path[1:]); err != nil {
		fmt.Println("error")
	}
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
