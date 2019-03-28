package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("rocketleague_dynamic_range.gohtml"))
}

func main() {
	carTypes := []string{"Merc", "Fast one", "Batman car"}

	err := tpl.Execute(os.Stdout, carTypes)
	if err != nil {
		log.Fatalln(err)
	}
}
