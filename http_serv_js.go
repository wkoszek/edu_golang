package main

import (
	"net/http"
	"net/http/httputil"

	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"

	//"time"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", handleTraffic)
	log.Fatal(http.ListenAndServe(":7125", nil))
}

func handleTraffic(out http.ResponseWriter, in *http.Request) {
	dbStr := "user=ubuntu dbname=ubuntu host=db sslmode=disable"
	hostStr, _ := os.Hostname()
	if strings.Contains(hostStr, "wkoszek") {
		dbStr = "user=ubuntu dbname=ubuntu host=localhost sslmode=disable"
	}

	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select mime_out,js_out from tags_get_many2(NULL)")
	//rows, err := db.Query("select mime,js from tags_get_many(ARRAY[1])")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if err != nil {
		fmt.Println("error!")
	} else {
		var (
			mime string
			js   string
		)

		for rows.Next() {
			err := rows.Scan(&mime, &js)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s %s\n", mime, js)
		}
	}

	reqStr, _ := httputil.DumpRequest(in, true)
	fmt.Println(string(reqStr))
}
