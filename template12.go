package main

import (
	"html/template"
	"log"
	"os"
)

const templateFile = "t12.tmpl"

var templates = template.Must(template.ParseFiles(templateFile))

func main() {
	log.Printf("template11")

	m := map[string]string{
		"first":  "something first",
		"second": "something second",
		"third":  "something third",
		"asd":    "asd",
	}

	err := templates.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
