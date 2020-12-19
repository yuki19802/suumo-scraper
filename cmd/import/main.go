package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/elastic/go-elasticsearch/v7"

	"github.com/yuki19802/suumo-scraper/internal/esdata"
	"github.com/yuki19802/suumo-scraper/internal/suumo"
)

func main() {
	///////////////////////////////////////////////////////////////////////////
	// Prepare ES
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)
	res.Body.Close()

	///////////////////////////////////////////////////////////////////////////
	// Do the importing
	importer := esdata.NewImporter(es)

	err = importer.DeleteAllDataIndices()

	if err != nil {
		log.Fatal("importer.DeleteAllDataIndices:", err)
	}

	ctx := context.TODO()

	err = filepath.Walk("./data", func(path string, info os.FileInfo, err error) error {
		log.Println("Loading", path)
		f, err := os.Open(path)

		if err != nil {
			return fmt.Errorf("os.Open: %w", err)
		}

		decoder := json.NewDecoder(f)

		var listings []suumo.Listing
		decoder.Decode(&listings)
		err = importer.ImportListings(ctx, listings)

		if err != nil {
			return fmt.Errorf("importer.ImportListings: %w", err)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
