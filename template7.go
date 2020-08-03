package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

const templateFile = "t7.tmpl"

var templates = template.Must(template.ParseFiles(templateFile, "head.tmpl"))

type Sample struct {
	Title  string
	Author string
}

type User struct {
	Username  string
	ItemsList []Sample
}

func main() {
	log.Printf("template1")

	isEmpty := flag.Bool("empty", false, "should we pass empty list")

	flag.Parse()

	log.Printf("bool:%t", *isEmpty)

	u := &User{
		Username: "admin",
		ItemsList: []Sample{
			Sample{"first", "Adam Koszek"},
			Sample{"second", "Bob"},
			Sample{"third", "Dan"},
		},
	}

	if *isEmpty {
		u = &User{}
	}

	err := templates.ExecuteTemplate(os.Stdout, templateFile, u)
	if err != nil {
		panic(err)
	}
}
