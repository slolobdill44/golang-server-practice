package main

import (
	"io"
	"net/http"
	"os"
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
	f, err := os.Open("11tigerpirate.jpg")
	if err != nil {
		http.Error(writer, "didn't find that file", 404)
		return
	}

	defer f.Close()

	io.Copy(writer, f)
}
