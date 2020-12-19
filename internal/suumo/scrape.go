package suumo

import (
	"fmt"
	"log"

	colly "github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

const (
	urlTokyo      = "https://suumo.jp/chintai/tokyo/city/"
	urlTokyo_sell = "https://suumo.jp/ms/chuko/aichi/city/"
)

func ListWards() ([]Ward, error) {
	wards := make([]Ward, 0)
	c := colly.NewCollector()

	c.OnHTML("#js-areaSelectForm", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, e *colly.HTMLElement) {
			name := e.ChildText("label span:first-of-type")
			code := e.ChildAttr("input[type=\"checkbox\"]", "value")
			log.Println(name, code)
			if code != "" {
				wards = append(wards, Ward{
					Name: name,
					Code: code,
				})
			}
		})
		log.Println(wards)
	})

	err := c.Visit(urlTokyo)

	if err != nil {
		return nil, err
	}

	return wards, nil
}

func ListWards2() ([]Ward, error) {
	wards := make([]Ward, 0)
	// c := colly.NewCollector()
	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	c.OnHTML("#js-areaSelectForm", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, e *colly.HTMLElement) {
			name := e.ChildText("label:not(span)")
			code := e.ChildAttr("input[type=\"checkbox\"]", "value")
			log.Println(name, code)
			if code != "" {
				wards = append(wards, Ward{
					Name: name,
					Code: code,
				})
			}
		})
		log.Println(wards)
	})

	err := c.Visit(urlTokyo_sell)

	if err != nil {
		return nil, err
	}

	return wards, nil
}

func WardListings2(ward Ward) ([]Listing, error) {
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

	//.property_unit-infoクラスの中の、dottable--cassetteクラスを走査する(この中はdivが色々定義)
	c.OnHTML("div.dottable--cassette", func(div *colly.HTMLElement) {
		name := div.ChildText("div.dottable-line:first-of-type dd")
		rawPrice := div.ChildText("div.dottable-line:nth-of-type(2) dd")
		// parsedPrice, err := extractPriceYen(rawPrice)
		// if err != nil {
		// 	panic(fmt.Sprintf("bad price: %q %v", rawPrice, err))
		// }

		neighborhood := div.ChildText("div.dottable-line:nth-of-type(3) dl:first-of-type dd")
		dist_to_station := div.ChildText("div.dottable-line:nth-of-type(3) dl:nth-of-type(2) dd")
		rawSquareMeters := div.ChildText("div.dottable-line:nth-of-type(4) td:first-of-type dd")
		// parsedSquareMeters, err := extractSquareMeters(rawSquareMeters)

		// if err != nil {
		// 	panic(fmt.Sprintf("bad square meters: %q %v", rawSquareMeters, err))
		// }

		rawYears := div.ChildText("div.dottable-line:nth-of-type(5) td:nth-of-type(2) dd")

		listing := Listing{
			Title:           name,
			Price:           rawPrice,
			Neighborhood:    neighborhood,
			Dist_to_station: dist_to_station,
			SquareMeters:    rawSquareMeters,
			AgeYears:        rawYears,
			Ward:            ward,
		}

		listings = append(listings, listing)
	})

	url_sell := fmt.Sprintf("https://suumo.jp/jj/bukken/ichiran/JJ010FJ001/?ar=050&bs=011&ta=23&sc=%s&kb=1&kt=9999999&mb=0&mt=9999999&ekTjCd=&ekTjNm=&tj=0&cnb=0&cn=9999999&pn=2", ward.Code)
	err := c.Visit(url_sell)
	if err != nil {
		return nil, err
	}
	return listings, nil
}
