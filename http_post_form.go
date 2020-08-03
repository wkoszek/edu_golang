package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	formData := url.Values{}
	formData.Set("value1", "samplevalue1")
	formData.Set("x", "1024")
	formData.Set("y", "768")
	resp, err := http.PostForm("http://localhost:7125/", formData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dbgResp, _ := httputil.DumpResponse(resp, false)
	fmt.Println(string(dbgResp))
}
