package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("rocketleague1.gohtml"))
	fmt.Println("hello, world")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "rocketleague1.gohtml", 4)
	if err != nil {
		log.Fatalln(err)
	}
}
