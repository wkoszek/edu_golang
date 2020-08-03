// Wojciech Adam Koszek
package main

import (
	//	"encoding/json"
	//	"crypto/md5"
	"flag"
	"fmt"
	"os"
)

const (
	g_author   = "Wojciech Adam Koszek"
	g_tagline  = "Nothing special"
	g_twitter  = "wkoszek"
	g_linkedin = "wkoszek"
)

var cmdArgVerbose = flag.Int("verbose", 0, "Turn the verbose logging")
var cmdArgOutputDir = flag.String("outputdir", "./_", "Output directory")

func parseArgs() []string {
	flag.Parse()
	//	cmdArg= flag.Int("flagname", 1234, "help message for flagname")
	//	flag.IntVar(&cmdArgVerbose, "verbose", defaultGopher, usage)
	return flag.Args()
}

type BlogPost struct {
	rawContent []byte
	tags       []string

	fileName string `json:"filename" yaml:"filename"`
	title    string `json:"title"    yaml:"title"`
	author   string `json:"author"   yaml:"author"`
}

func processFile(fileName string) (*BlogPost, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, nil
	}
	var fileSize = stat.Size()
	var fileContent = make([]byte, fileSize)
	var fileBytesRead = 0

	fileBytesRead, err = file.Read(fileContent)
	if err != nil {
		fmt.Println("nothing here")
		return nil, nil
	}
	file.Close()

	var strFileContent = string(fileContent)

	fmt.Println("File size is ", fileSize, "read: ", fileBytesRead)
	fmt.Println("File stuff is: %s", strFileContent)

	return &BlogPost{author: "Wojciech Adam Koszek"}, nil
}

func main() {
	var args = parseArgs()

	fmt.Println("============== args =================")
	fmt.Println(args)

	progName := os.Args[0]
	fmt.Println(progName)

	argsWithoutProg := os.Args[1:]

	var allBlogPosts [](*BlogPost)
	for _, arg := range argsWithoutProg {
		fmt.Println(arg)
		var b, _ = processFile(arg)
		allBlogPosts = append(allBlogPosts, b)
	}

	fmt.Println("----------------------")
	fmt.Println(allBlogPosts)
}
