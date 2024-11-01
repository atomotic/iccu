package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/atomotic/iccu/sbn"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Record struct {
	ID         string   `json:"bid"`
	IDLink     string   `json:"id"`
	IDMOL      string   `json:"idmanus"`
	Title      string   `json:"title"`
	IIIF       []string `json:"iiif,omitempty"`
	Link       []string `json:"link,omitempty"`
	DocType    string   `json:"type,omitempty"`
	MatType    []string `json:"material,omitempty"`
	Thumbnails []string `json:"thumbnails,omitempty"`
	StartDate  int      `json:"start_date,omitempty"`
	EndDate    int      `json:"end_date,omitempty"`
}

func main() {

	db, err := sqlx.Open("sqlite3", "sbn-metadata.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Queryx("select doc from sbn")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var jsonData string
		err := rows.Scan(&jsonData)
		if err != nil {
			log.Fatal(err)
		}

		var doc sbn.Doc
		err = json.Unmarshal([]byte(jsonData), &doc)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %v", err)
			continue
		}

		record := Record{
			IDLink:     doc.BidLink(),
			ID:         doc.Bid(),
			IDMOL:      doc.IDMol(),
			Title:      doc.Title(),
			IIIF:       doc.Manifests(),
			Link:       doc.ExternalLink(),
			DocType:    doc.Type(),
			MatType:    doc.Material(),
			Thumbnails: doc.Thumbnails(),
			StartDate:  doc.StartDate(),
			EndDate:    doc.EndDate(),
		}

		marshaledRecord, err := json.Marshal(record)
		if err != nil {
			continue
		}

		fmt.Println(string(marshaledRecord))

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
