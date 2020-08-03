package main

import (
	"net/http"
	"net/http/httputil"

	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", handleTraffic)
	log.Fatal(http.ListenAndServe("localhost:7125", nil))
}

func handleTraffic(w http.ResponseWriter, r *http.Request) {
	keys, err := r.URL.Query()["code"]
	if len(keys) < 1 || err != true {
		log.Println("error !")
	}
	log.Println("keys => %v", keys)

	key := keys[0]
	log.Println("key => %v", key)

	reqStr, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(reqStr))
}
