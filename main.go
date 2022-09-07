package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
func main() {
	fmt.Println("Server start")
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}
