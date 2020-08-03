package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

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

	pageStr := flag.String("page", "about.tmpl", "template name")

	flag.Parse()

	log.Println(*pageStr)

	var templates = template.Must(template.ParseFiles("base.tmpl", *pageStr))

	log.Printf("page:%s", pageStr)

	err := templates.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
