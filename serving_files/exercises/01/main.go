package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/11tigerpirate.jpg", pirate)
	http.ListenAndServe(":8081", nil)
}

func foo(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "foo ran")
}

func dog(writer http.ResponseWriter, req *http.Request) {
	http.ServeFile(writer, req, "dog.gohtml")

}

func pirate(writer http.ResponseWriter, req *http.Request) {
	http.ServeFile(writer, req, "11tigerpirate.jpg")
}
