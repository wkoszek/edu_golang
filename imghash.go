package main

import (
	"fmt"
	"log"
	"os"

	"github.com/carlogit/phash"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: %s <ImageFileA> <ImageFileB>\n", os.Args[0])
	}

	var w = "wojtek"
	fmt.Println(w[len(w)-3:])
	os.Exit(0)
	a := hash(os.Args[1])
	b := hash(os.Args[2])

	fmt.Println(a)
	fmt.Println(b)

	distance := phash.GetDistance(a, b)

	fmt.Printf("distance: %d\n", distance)
}

// hash returns a phash of the image
func hash(filename string) string {
	img, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	ahash, err := phash.GetHash(img)
	if err != nil {
		log.Fatal(err)
	}
	return ahash
}
