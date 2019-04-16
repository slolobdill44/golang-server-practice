package main

import (
	"net/http"
)

func main() {

	//server := http.FileServer(http.Dir("ind"))

	http.ListenAndServe(":8081", http.FileServer(http.Dir(".")))

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("/public"))))

}
