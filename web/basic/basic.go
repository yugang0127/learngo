package main

import (
	"log"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main()  {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}
