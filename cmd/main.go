package main

import (
	"fmt"
	"os"

	"github.com/evertras/suumo-scraper/internal/suumo"
)

func main() {
	listings, err := suumo.WardListings("13101")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Total hits:", len(listings))

	for _, l := range listings {
		if l.Title != "" {
			fmt.Printf("%+v\n", l)
		}
	}
}
