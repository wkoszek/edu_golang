package main;

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"time"
	"strings"
	"os"
)

func main() {
	time.Sleep(time.Second * 5)

	hostStr, _ := os.Hostname()

	dbStr := "user=ubuntu dbname=ubuntu host=db sslmode=disable"

	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select mime_out,js_out from tags_get_many2(NULL)")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if err != nil {
		fmt.Println("error!")
	} else {
		var (
			mime	string
			js	string
		)

		for rows.Next() {
			err := rows.Scan(&mime, &js)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s %s\n", mime, js)
		}
	}
}
