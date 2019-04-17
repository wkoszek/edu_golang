package main;

import (
	"fmt"
)

func funcMaker(someMagicConst int) func(string) {
	return func(a string) {
		fmt.Println(someMagicConst, a)
	}
}

func writeString(strIn string) {
	fmt.Println(strIn)
}

func main() {
	db := dbGet()
	defer db.close()


	http.HandleFunc(wrapGet(wrapDb("/asd", tagsGet)))

	c.Get(dbProvideHttp("/tags", getAsd))

	f := funcMaker(123)
	f("asodj")
}
