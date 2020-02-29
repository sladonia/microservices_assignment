package domains

type Port struct {
	Abbreviation string    `json:"abbreviation"`
	Name         string    `json:"name"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	Alias        []string  `json:"alias"`
	Regions      []string  `json:"regions"`
	Coordinates  []float64 `json:"coordinates"`
	Province     string    `json:"province"`
	Timezone     string    `json:"timezone"`
	Unlocs       []string  `json:"unlocs"`
	Code         string    `json:"code"`
}

type ImportResponse struct {
	NumberInserted int32 `json:"number_inserted"`
	NumberUpdated  int32 `json:"number_updated"`
}
