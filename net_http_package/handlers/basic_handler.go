package main

import (
	"fmt"
	"net/http"
)

type content int

func (w content) ServeHTTP(m http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(m, "Any code you want in this func")
}

func main() {
	var what content
	http.ListenAndServe(":8081", what)
}
