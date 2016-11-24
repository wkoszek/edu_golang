package main

import (
    "fmt"
    "time"
    "os"
    "log"
    "path/filepath"
)

func main() {


	if (len(os.Args) != 2) {
		log.Fatal("bleh")
	}
	searchDir := os.Args[1]

	timeBefore := time.Now()
	fileList := []string{}
		err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	timeAfter := time.Now()

	fmt.Println("============= Time difference ============")
	diff := timeAfter.Sub(timeBefore)
	fmt.Println(diff)

	fmt.Println("=========== Will wait here ==============")
	time.Sleep(3000*time.Second)

	if (false) {
		for _, file := range fileList {
			fmt.Println(file)
		}
		fmt.Println(err)
	}
}
