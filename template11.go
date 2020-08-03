package main

import (
	"html/template"
	"log"
	"os"
)

const templateFile = "t11.tmpl"

var templates = template.Must(template.ParseFiles(templateFile))

func main() {
	log.Printf("template11")

	u := make([]int, 10)

	log.Print(len(u))

	u[0] = 88
	u[1] = 11
	u[2] = 22

	err := templates.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}
}
