package main

import (
	"html/template"
	"log"
	"os"
)

const templateFile = "t14.tmpl"

func rowColored(rowNum int) bool {
	//rowNum, err := strconv.Atoi(rowNumStr);
	//if err == nil {
	//	return false
	//}
	return rowNum%2 == 0
}

func main() {
	log.Printf("template14")

	m := []string{
		"first",
		"second",
		"third",
		"asd",
	}

	fs := template.FuncMap{
		"rowColored": rowColored,
	}

	t, err := template.New(templateFile).Funcs(fs).ParseFiles(templateFile)

	err = t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
