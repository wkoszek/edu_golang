package main;

import (
	"net/http"
	"net/http/httputil"

	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", handleTraffic)
	log.Fatal(http.ListenAndServe(":7125", nil))
}

func handleTraffic(out http.ResponseWriter, in *http.Request) {
	reqStr,  _ := httputil.DumpRequest(in, true)
	fmt.Println(string(reqStr))
}
