package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	req, err := http.NewRequest("PUT", "http://localhost:7125", strings.NewReader("data"))
	if err != nil {
		log.Fatal(err)
	}

	cli := &http.Client{}
	resp, err2 := cli.Do(req)
	if err2 != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dbgResp, _ := httputil.DumpResponse(resp, false)
	fmt.Println(string(dbgResp))
}
