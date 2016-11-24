package main

import (
	"gopkg.in/yaml.v2"
	"github.com/kr/pretty"
	"path/filepath"
	"io/ioutil"
	"runtime"
	"flag"
	"fmt"
)

type Article struct {
	title		string `yaml:"title"`
//	author		string
//	abstract	string
//	description	string // probably to nix -- same as abstract?
//	address		string
//	auth		string
//	read		string
//	layout		string
//	tags		[]string
//	categories	[]string	// probably to nix -- same as tags?
//	ads		[]string
//	allow_words	[]string
}

func main() {
	var cmd_yamlv2 =  flag.Bool("yamlv2",  false, "use yaml.v2")
	var cmd_verbose = flag.Bool("verbose", false, "use verbose")
	var cmd_gypsy =   flag.Bool("gypsy",   false, "gypsy")

	flag.Parse()

	var argGot = 0
	if (*cmd_yamlv2) {
		argGot += 1
	}
	if (*cmd_gypsy) {
		argGot += 1
	}
	fmt.Println("argGot: ", argGot)

	var yamlFileSet, err = filepath.Glob("/Users/wkoszek/go/w.yml")
	if (err != nil) {
		fmt.Println("error!")
	}
	var yamlFileName = yamlFileSet[0]
	fmt.Println(yamlFileName)
	yamlFile, err := ioutil.ReadFile(yamlFileName)
	fmt.Println(yamlFile)
	check(err)

	if (*cmd_verbose) {
		fmt.Println(string(yamlFile))
	}
	if (*cmd_yamlv2) {
		use_yamlv2(yamlFile, yamlFileName)
	}
}

func use_yamlv2(yamlFile []byte, filename string) {
	fmt.Println("========================================")
	fmt.Println(runtime.Caller(0))

	var article Article
	fmt.Printf("xxxxxxxxxxxx %#v xxxxxxxxxxxxx", yamlFile)

	var err = yaml.Unmarshal(yamlFile, &article)
	check(err)

	fmt.Println("--------------------")
	fmt.Printf("Description: %#v\n", article)
	fmt.Println("--------------------")
	fmt.Println("%# v", pretty.Formatter(article))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
