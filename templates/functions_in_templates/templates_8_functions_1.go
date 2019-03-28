package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("rocketcars_func.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	carTypes := map[string]string{"wide": "Merc", "speedy": "Octane", "premium": "Batman car"}

	err := tpl.ExecuteTemplate(os.Stdout, "rocketcars_func.gohtml", carTypes)
	if err != nil {
		log.Fatalln(err)
	}
}
