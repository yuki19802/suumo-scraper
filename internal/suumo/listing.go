package suumo

// type Listing struct {
// 	Title            string  `json:"title"`
// 	Neighborhood     string  `json:"neighborhood"`
// 	AgeYears         int     `json:"ageYears"`
// 	Floor            int     `json:"floor"`
// 	PricePerMonthYen int     `json:"pricePerMonthYen"`
// 	Layout           string  `json:"layout"`
// 	SquareMeters     float32 `json:"squareMeters"`
// 	Ward             Ward    `json:"ward"`
// }

type Listing struct {
	Title           string `json:"title"`
	Price           string `json:"price"`
	Neighborhood    string `json:"neighborhood"`
	Dist_to_station string `json:"dist_to_station"`
	SquareMeters    string `json:"rawSquareMeters"`
	AgeYears        string `json:"rawYears"`
	Ward            Ward   `json:"ward"`
}
