package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	postBody := strings.NewReader("some body")
	resp, err := http.Post("http://localhost/", "text/plain", postBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dbgResp, _ := httputil.DumpResponse(resp, false)
	fmt.Println(string(dbgResp))
}
