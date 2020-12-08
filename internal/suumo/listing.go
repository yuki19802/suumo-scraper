package suumo

type Listing struct {
	Title            string  `json:"title"`
	Neighborhood     string  `json:"neighborhood"`
	AgeYears         int     `json:"ageYears"`
	Floor            int     `json:"floor"`
	PricePerMonthYen int     `json:"pricePerMonthYen"`
	Layout           string  `json:"layout"`
	SquareMeters     float32 `json:"squareMeters"`
	Ward             Ward    `json:"ward"`
}
