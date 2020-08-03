package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

var templates = template.Must(template.ParseFiles("t3.tmpl"))

type Sample struct {
	Title  string
	Author string
}

func main() {
	log.Printf("template1")

	isEmpty := flag.Bool("empty", false, "should we pass empty list")

	flag.Parse()

	log.Printf("bool:%t", *isEmpty)

	pages := []Sample{
		Sample{"first", "Adam Koszek"},
		Sample{"second", "Bob"},
		Sample{"third", "Dan"},
	}

	if *isEmpty {
		pages = []Sample{}
	}

	err := templates.ExecuteTemplate(os.Stdout, "t3.tmpl", pages)
	if err != nil {
		panic(err)
	}
}
