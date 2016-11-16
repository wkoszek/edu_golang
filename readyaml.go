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
	posts_all	[Max_posts_number]Post
)

func main() {
	var files []string

	files, err = filepath.Glob(Posts_path)
	failonerr(err)

	var p Post
	for pi, filename := range(files) {
		p = processfile(filename)
		posts_all[pi] = p
	}

	os.Exit(0)
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

	if (false) {
		fmt.Println(filename)
		fmt.Printf("--- t:\n%+v\n\n", p)
	}
	return p
}

func failonerr(err error) {
	if (err != nil) {
		log.Fatal(err)
	}
}

