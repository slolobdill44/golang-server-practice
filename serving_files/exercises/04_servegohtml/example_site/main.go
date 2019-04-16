package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	//server := http.FileServer(http.Dir("ind"))

	http.HandleFunc("/", home)

	http.ListenAndServe(":8081", nil)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("/public"))))

}

func home(writer http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("/templates/index.gohtml"))

	err := tpl.Execute(writer, nil)

	if err != nil {
		log.Fatal("problem parsing gohtml")
	}
}
