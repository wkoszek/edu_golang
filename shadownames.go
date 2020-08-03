package main

import (
	"fmt"
)

func main() {
	fmt.Println("something")
	i := 13
	if true {
		i := 17
		fmt.Printf("changing i! to %d\n", i)
	}
	fmt.Println(i)
}
