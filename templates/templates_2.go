package main 

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("goindex.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error in creating file", err)
	}

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}