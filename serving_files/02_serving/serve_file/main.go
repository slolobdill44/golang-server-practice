package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/pirate", servePirate)
	http.ListenAndServe(":8081", nil)
}

func homePage(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(writer, `
	<img src="11tigerpirate.jpg"></img>`)
}

func servePirate(writer http.ResponseWriter, req *http.Request) {
	http.ServeFile(writer, req, "11tigerpirate.jpg")
}
