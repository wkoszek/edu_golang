package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://localhost/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dbgResp, _ := httputil.DumpResponse(resp, false)
	fmt.Println(string(dbgResp))
}
