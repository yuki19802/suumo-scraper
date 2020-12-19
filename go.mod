module github.com/yuki19802/suumo-scraper

go 1.15

require (
	github.com/elastic/go-elasticsearch/v7 v7.5.1-0.20201026095746-17c910458d57
	github.com/gocolly/colly/v2 v2.1.0
)

//ローカルをimportする指定
replace github.com/yuki19802/suumo-scraper/internal/suumo => ../../internal/suumo
