package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/evertras/suumo-scraper/internal/suumo"
)

func main() {
	log.Println("Getting wards...")

	wards, err := suumo.ListWards()

	if err != nil {
		panic(err)
	}

	log.Printf("Found %d wards", len(wards))

	for _, ward := range wards {
		path := fmt.Sprintf("data/%s.json", ward.Code)

		if _, err := os.Stat(path); !os.IsNotExist(err) {
			log.Printf("Data already exists for %s (%s), skipping...", ward.Name, ward.Code)
			continue
		}

		log.Printf("Fetching %s (%s)", ward.Name, ward.Code)

		listings, err := suumo.WardListings(ward)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Total hits:", len(listings))

		file, err := os.Create(path)

		if err != nil {
			log.Fatal("Failed to open file to write:", err)
		}

		encoder := json.NewEncoder(file)

		err = encoder.Encode(listings)

		if err != nil {
			log.Fatal("Failed to encode JSON to file:", err)
		}
	}
}
