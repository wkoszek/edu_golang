package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

const templateFile = "t8.tmpl"

var templates = template.Must(template.ParseFiles(templateFile, "head.tmpl", "mosaic.tmpl"))

type Sample struct {
	Title  string
	Author string
	Nums   []int
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
			Sample{"first", "Adam Koszek", []int{1, 2, 3}},
			Sample{"second", "Bob", []int{10, 13, 28}},
			Sample{"third", "Dan", []int{23, 55, 84}},
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
