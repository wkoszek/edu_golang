package main

import (
	"html/template"
	"log"
	"os"
)

const templateFile = "t13.tmpl"

func charCount(s string) int {
	return len(s)
}

func main() {
	log.Printf("template13")

	m := map[string]string{
		"first":  "something first",
		"second": "something second",
		"third":  "something third",
		"asd":    "asd",
	}

	fs := template.FuncMap{
		"charcount": charCount,
	}

	t, err := template.New(templateFile).Funcs(fs).ParseFiles(templateFile)

	err = t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
