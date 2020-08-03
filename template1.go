package main

import (
	"html/template"
	"log"
	"os"
)

var templates = template.Must(template.ParseFiles("t1.tmpl"))

type Sample struct {
	Title string
}

func main() {
	log.Printf("template1")

	p := &Sample{"sample"}

	err := templates.ExecuteTemplate(os.Stdout, "t1.tmpl", p)
	if err != nil {
		panic(err)
	}
}
