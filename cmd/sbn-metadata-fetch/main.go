package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/atomotic/iccu/client"
	"github.com/carlmjohnson/requests"
	_ "github.com/mattn/go-sqlite3"
)

var base = "https://api.iccu.sbn.it/sbn/1.0.0/search?format=json&detail=full&page-size=500&presenza_digitale=Y"
var c *http.Client

func main() {
	client.New(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
	)
	ctx := context.Background()
	c = client.GetClient()

	db, err := sql.Open("sqlite3", "sbn-metadata.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pragmas := []string{
		"PRAGMA journal_mode = WAL",
		"PRAGMA synchronous = NORMAL",
		"PRAGMA journal_size_limit = 67108864",
		"PRAGMA mmap_size = 134217728",
		"PRAGMA cache_size = 2000",
		"PRAGMA busy_timeout = 5000",
	}

	for _, pragma := range pragmas {
		_, err := db.Exec(pragma)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS sbn (
		bid TEXT GENERATED ALWAYS AS (json_extract(doc, '$.unimarc.fields[1].003')) VIRTUAL, 
		doc json
	); 
	CREATE INDEX IF NOT EXISTS bid_idx on sbn(bid);`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM sbn")
	if err != nil {
		log.Fatal(err)
	}

	pageSize := 500
	startAt := 0
	processed := 0

	for {
		urlx := fmt.Sprintf("%s&start-at=%d", base, startAt)

		var response client.FullResponse
		err := requests.URL(urlx).Client(c).ToJSON(&response).Header("Accept-Encoding", "gzip").Fetch(ctx)

		if err != nil {
			fmt.Println("Error fetching data", err)
			startAt += pageSize
			continue
		}

		total := response.Response.NumFound

		for _, doc := range response.Response.Docs {
			j, _ := json.Marshal(doc)
			_, err := db.Exec("INSERT INTO sbn(doc) values(?)", j)
			if err != nil {
				log.Fatal(err)
			}
			processed++
		}

		if processed >= total {
			break
		}

		startAt = processed
		fmt.Printf("\r%d / %d", processed, total)
	}

}
