package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	req, err := http.NewRequest("DELETE", "http://localhost:7125", nil)
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
