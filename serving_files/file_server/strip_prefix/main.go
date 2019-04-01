package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/pirate", pirate)
	http.ListenAndServe(":8081", nil)
}

func pirate(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf8")
	io.WriteString(writer, `
		<img src="/resources/11tigerpirate.jpg"></img>`)
}
