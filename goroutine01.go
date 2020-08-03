package main

import (
	"fmt"
	"time"
)

func gor(id int, t *time.Time) {
	//	b := time.Now()
	*t = time.Now()

	time.Sleep(time.Millisecond * 500)

	//	a := time.Now()
	//	fmt.Printf("%+v %+v\n", a, b)
	//	*t = b
}

func main() {
	fmt.Println("starting")

	const how_many_threads = 10

	times := make([]time.Time, how_many_threads)

	fmt.Println(times)

	for i := 0; i < how_many_threads; i++ {
		go gor(i, &times[i])
	}

	fmt.Println(times)

	fmt.Println("========== list ===========")
	for i := 0; i < how_many_threads; i++ {
		fmt.Println(times[i])
	}
}
