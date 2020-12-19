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

// func WardListings(ward Ward) ([]Listing, error) {
// 	listings := make([]Listing, 0)
// 	c := colly.NewCollector()

// 	c.OnRequest(func(r *colly.Request) {
// 		log.Printf("Visiting page %s - %s", r.URL.Query().Get("page"), r.URL)
// 	})

// 	c.OnHTML(".pagination", func(e *colly.HTMLElement) {
// 		e.ForEach("p.pagination-parts:last-of-type a", func(i int, a *colly.HTMLElement) {
// 			if a.Text == "次へ" {
// 				a.Request.Visit(a.Attr("href"))
// 			}
// 		})
// 	})
// 	//.l-cassetteitemクラスの中のliを走査する
// 	c.OnHTML(".l-cassetteitem > li", func(li *colly.HTMLElement) {
// 		//liの中の.cassetteitem_content-titleのテキストを取ってくる（クラスが一位に決まる場合はこれでOK）
// 		name := li.ChildText(".cassetteitem_content-title")
// 		neighborhood := li.ChildText(".cassetteitem_detail-col1")
// 		//liの中のcassetteitem_detail-col3のうち、最初のdivテキストを取ってくる。
// 		yearsRaw := li.ChildText(".cassetteitem_detail-col3 > div:first-of-type")

// 		parsedYears, err := extractAgeYears(yearsRaw)

// 		if err != nil {
// 			panic(fmt.Sprint("bad year:", yearsRaw, err))
// 		}
//        //liの中の<tr class="js-cassette_link">を全て見ていく。
// 		li.ForEach("tr.js-cassette_link", func(i int, tr *colly.HTMLElement) {
// 			//trの中の3番目のtd要素を持ってくる（cssの書き方そのもの）
// 			rawFloor := tr.ChildText("td:nth-child(3)")

// 			parsedFloor, err := extractFloor(rawFloor)

// 			if err != nil {
// 				panic(fmt.Sprint("bad floor:", rawFloor, err))
// 			}
//             //trの中の4番目のtd要素のうち、cassetteitem_price--rentクラスを持ってくる（cssの書き方そのもの）
// 			rawPrice := tr.ChildText("td:nth-child(4) .cassetteitem_price--rent")

// 			parsedPrice, err := extractPriceYen(rawPrice)

// 			if err != nil {
// 				panic(fmt.Sprintf("bad price: %q %v", rawPrice, err))
// 			}

// 			layout := tr.ChildText(".cassetteitem_madori")

// 			rawSquareMeters := tr.ChildText(".cassetteitem_menseki")

// 			parsedSquareMeters, err := extractSquareMeters(rawSquareMeters)

// 			if err != nil {
// 				panic(fmt.Sprintf("bad square meters: %q %v", rawSquareMeters, err))
// 			}

// 			listing := Listing{
// 				Title:            name,
// 				Neighborhood:     neighborhood,
// 				AgeYears:         parsedYears,
// 				Floor:            parsedFloor,
// 				PricePerMonthYen: parsedPrice,
// 				Layout:           layout,
// 				SquareMeters:     parsedSquareMeters,
// 				Ward:             ward,
// 			}

// 			listings = append(listings, listing)
// 		})
// 	})

// 	url := fmt.Sprintf("https://suumo.jp/jj/chintai/ichiran/FR301FC001/?ar=030&bs=040&ta=13&sc=%s&cb=0.0&ct=9999999&mb=0&mt=9999999&et=9999999&cn=9999999&shkr1=03&shkr2=03&shkr3=03&shkr4=03&sngz=&po1=25&pc=10&page=1", ward.Code)
// 	// url := fmt.Sprintf("https://suumo.jp/jj/bukken/ichiran/JJ012FC001/?ar=030&bs=011&cn=9999999&cnb=0&ekTjCd=&ekTjNm=&kb=1&kt=9999999&mb=0&mt=9999999&sc=%s&ta=13&tj=0&po=0&pj=1&pc=30&page=1&pc=2&pn=1", ward.Code)

// 	err := c.Visit(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return listings, nil
// }

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
		// neighborhood := div.ChildText("dd.dottable-vm")
		listing := Listing{
			Title: name,
			Price: rawPrice,
			// Neighborhood: neighborhood,
			Ward: ward,
		}

		listings = append(listings, listing)
	})

	url_sell := fmt.Sprintf("https://suumo.jp/jj/bukken/ichiran/JJ012FC001/?ar=030&bs=011&cn=9999999&cnb=0&ekTjCd=&ekTjNm=&kb=1&kt=9999999&mb=0&mt=9999999&sc=%s&ta=13&tj=0&po=0&pj=1&pc=30&page=1&pc=2&pn=1", ward.Code)
	err := c.Visit(url_sell)
	if err != nil {
		return nil, err
	}
	return listings, nil
}
