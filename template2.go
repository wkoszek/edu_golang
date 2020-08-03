package main

import (
	"html/template"
	"log"
	"os"
)

var templates = template.Must(template.ParseFiles("t2.tmpl"))

type Sample struct {
	Title  string
	Author string
}

func main() {
	log.Printf("template1")

	p := &Sample{"sample", "Adam Koszek"}

	err := templates.ExecuteTemplate(os.Stdout, "t2.tmpl", p)
	if err != nil {
		panic(err)
	}
}
