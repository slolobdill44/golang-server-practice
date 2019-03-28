package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func pourBeer(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "psssssssh fttttt ok ready to drink")
}

func drinkBeer(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "glug glug yeah fresh")
}

func me(resp http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("template.gohtml")

	if err != nil {
		log.Fatalln("There was an error parsing that")
	}

	err = tpl.ExecuteTemplate(resp, "template.gohtml", "Adrian")

	if err != nil {
		log.Fatalln("There was an error executing that")
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(pourBeer))
	http.Handle("/drink/", http.HandlerFunc(drinkBeer))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8081", nil)
}
