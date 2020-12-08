package suumo

import (
	"fmt"
	"log"

	colly "github.com/gocolly/colly/v2"
)

const (
	urlTokyo = "https://suumo.jp/chintai/tokyo/city/"
)

func ListWards() ([]Ward, error) {
	wards := make([]Ward, 0)

	c := colly.NewCollector()

	c.OnHTML("#js-areaSelectForm", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, e *colly.HTMLElement) {
			name := e.ChildText("label span:first-of-type")
			code := e.ChildAttr("input[type=\"checkbox\"]", "value")

			if code != "" {
				wards = append(wards, Ward{
					Name: name,
					Code: code,
				})
			}
		})
	})

	err := c.Visit(urlTokyo)

	if err != nil {
		return nil, err
	}

	return wards, nil
}

func WardListings(ward Ward) ([]Listing, error) {
	listings := make([]Listing, 0)
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting page %s - %s", r.URL.Query().Get("page"), r.URL)
	})

	c.OnHTML(".pagination", func(e *colly.HTMLElement) {
		e.ForEach("p.pagination-parts:last-of-type a", func(i int, a *colly.HTMLElement) {
			if a.Text == "次へ" {
				a.Request.Visit(a.Attr("href"))
			}
		})
	})

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

			rawPrice := tr.ChildText("td:nth-child(4) .cassetteitem_price--rent")

			parsedPrice, err := extractPriceYen(rawPrice)

			if err != nil {
				panic(fmt.Sprintf("bad price: %q %v", rawPrice, err))
			}

			layout := tr.ChildText(".cassetteitem_madori")

			rawSquareMeters := tr.ChildText(".cassetteitem_menseki")

			parsedSquareMeters, err := extractSquareMeters(rawSquareMeters)

			if err != nil {
				panic(fmt.Sprintf("bad square meters: %q %v", rawSquareMeters, err))
			}

			listing := Listing{
				Title:            name,
				Neighborhood:     neighborhood,
				AgeYears:         parsedYears,
				Floor:            parsedFloor,
				PricePerMonthYen: parsedPrice,
				Layout:           layout,
				SquareMeters:     parsedSquareMeters,
				Ward:             ward,
			}

			listings = append(listings, listing)
		})
	})

	url := fmt.Sprintf("https://suumo.jp/jj/chintai/ichiran/FR301FC001/?ar=030&bs=040&ta=13&sc=%s&cb=0.0&ct=9999999&mb=0&mt=9999999&et=9999999&cn=9999999&shkr1=03&shkr2=03&shkr3=03&shkr4=03&sngz=&po1=25&pc=10&page=1", ward.Code)

	err := c.Visit(url)

	if err != nil {
		return nil, err
	}

	return listings, nil
}
