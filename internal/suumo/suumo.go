package suumo

import (
	"fmt"

	colly "github.com/gocolly/colly/v2"
)

// Chiyoda:  https://suumo.jp/jj/chintai/ichiran/FR301FC001/?ar=030&bs=040&ta=13&sc=13101&cb=0.0&ct=9999999&et=9999999&cn=9999999&mb=0&mt=9999999&shkr1=03&shkr2=03&shkr3=03&shkr4=03&fw2=&srch_navi=1

const (
	urlTokyo = "https://suumo.jp/chintai/tokyo/city/"
)

func ListWards() error {
	c := colly.NewCollector()

	c.OnHTML("#js-areaSelectForm", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, e *colly.HTMLElement) {
			name := e.ChildText("label span:first-of-type")
			code := e.ChildAttr("input[type=\"checkbox\"]", "value")

			fmt.Println(i, name, code)
		})
	})

	return c.Visit(urlTokyo)
}

func WardListings(wardCode string) ([]Listing, error) {
	listings := make([]Listing, 0)
	c := colly.NewCollector()
	url := fmt.Sprintf("https://suumo.jp/jj/chintai/ichiran/FR301FC001/?ar=030&bs=040&ta=13&sc=%s&cb=0.0&ct=9999999&et=9999999&cn=9999999&mb=0&mt=9999999&shkr1=03&shkr2=03&shkr3=03&shkr4=03&fw2=&srch_navi=1", wardCode)

	fmt.Println(url)
	c.OnHTML(".l-cassetteitem > li", func(li *colly.HTMLElement) {
		name := li.ChildText(".cassetteitem_content-title")
		neighborhood := li.ChildText(".cassetteitem_detail-col1")
		yearsRaw := li.ChildText(".cassetteitem_detail-col3 > div:first-of-type")

		parsedYears, err := extractAgeYears(yearsRaw)

		if err != nil {
			panic(fmt.Sprint("bad year:", yearsRaw, err))
		}

		li.ForEach("tr.js-cassette_link", func(i int, tr *colly.HTMLElement) {
			rawFloor := tr.ChildText("td:nth-child(3)")

			parsedFloor, err := extractFloor(rawFloor)

			if err != nil {
				panic(fmt.Sprint("bad floor:", rawFloor, err))
			}

			listing := Listing{
				Title: name,
				Neighborhood: neighborhood,
				AgeYears: parsedYears,
				Floor: parsedFloor,
			}

			listings = append(listings, listing)
		})
	})

	err := c.Visit(url)

	if err != nil {
		return nil, err
	}

	return listings, nil
}
