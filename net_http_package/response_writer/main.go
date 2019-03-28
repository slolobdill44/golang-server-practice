package main

import (
	"fmt"
	"net/http"
)

type mywebsite int

func (m mywebsite) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Very-Cool_Header", "8)")
	writer.Header().Set("Other-Cool-Header", "OOOoooohh yeahhh")
	fmt.Fprintf(writer, "<h1>sup peeps lets drink orange juice</h1>")
}

func main() {
	var site mywebsite
	http.ListenAndServe(":8081", site)
}
