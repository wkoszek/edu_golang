package main;

import (
	"net/http"
	"net/http/httputil"
	"strings"
	"fmt"
	"log"
)

func main() {
	postBody := strings.NewReader("some body")
	resp, err := http.Post("http://localhost:7125/", "text/plain", postBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dbgResp, _ := httputil.DumpResponse(resp, false)
	fmt.Println(string(dbgResp))
}
