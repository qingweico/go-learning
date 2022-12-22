package main

import (
	"fmt"
	"io"
	"net/http"
)

func simExample(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello World!")
	if err != nil {
		return
	}
	fmt.Println(r)
}

func main() {
	http.HandleFunc("/", simExample)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		return
	}
}
