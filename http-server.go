package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/hello", helloworld)

	http.ListenAndServe(":8090", nil)
}
