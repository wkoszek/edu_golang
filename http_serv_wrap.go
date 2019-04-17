package main;

import (
	"net/http"
	"net/http/httputil"

	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", makeHandler(handleTraffic, 123))

	log.Fatal(http.ListenAndServe("localhost:7125", nil))
}

func makeHandler(handler func(http.ResponseWriter, *http.Request, int), mag int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, mag)
	}
}

func handleTraffic(out http.ResponseWriter, in *http.Request, magic int) {
	reqStr,  _ := httputil.DumpRequest(in, true)
	fmt.Println(string(reqStr))
}
