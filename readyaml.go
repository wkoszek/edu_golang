package main

import (
	"gopkg.in/yaml.v2"
	"path/filepath"
	"fmt"
	"log"
	"os"
)

type Post struct {
	Title	   string	`yaml:"title"`
	Author	   string	`yaml:"author"`
	Abstract   string	`yaml:"abstract"`
	Address	   string	`yaml:"address"`
	Auth	   string	`yaml:"auth"`
	Read	   string	`yaml:"read"`
	Layout	   string	`yaml:"layout"`
	Tags	 []string	`yaml:"tags"`
	Ads	 []string	`yaml:"ads"`
	Body	 []string
}

const (
	Posts_path = "/Users/wk/r/me/source/blog/*.md"
	Max_posts_number = 10000
)

var (
	err		  error
	posts_all	[]Post
)

func main() {
	var files []string
	var p Post

	files, err = filepath.Glob(Posts_path)
	failonerr(err)

	for _, filename := range(files) {
		p = processfile(filename)
		posts_all = append(posts_all, p)
	}

	renderAll(posts_all)
	os.Exit(0)
}

func failonerr(err error) {
	if (err != nil) {
		log.Fatal(err)
	}
}

func processfile(filename string) Post {
	var file *os.File
	var st os.FileInfo
	var filebody []byte

	file, err = os.Open(filename)
	failonerr(err)

	st, err = file.Stat()
	failonerr(err)

	filebody = make([]byte, st.Size())
	_, err = file.Read(filebody)
	err = file.Close()
	failonerr(err)

	return readyaml(filename, filebody)
}

func
readyaml(filename string, body []byte) Post {
	p := Post{}
	err := yaml.Unmarshal(body, &p)
	failonerr(err)

	if (true) {
		fmt.Println(filename)
		fmt.Printf("--- t:\n%+v\n\n", p)
	}
	return p
}

func renderAll(allPosts []Post) bool {
//	tmpl, err := template.New("main").Parse("{{.Count}} items are made of {{.Material}}")
//	if err != nil { panic(err) }
//	err = tmpl.Execute(os.Stdout, sweaters)
//	if err != nil { panic(err) }
	return true
}
