package domains

import "client_api/src/portpb"

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

func PortFromPBObject(p *portpb.Port) *Port {
	port := &Port{
		Abbreviation: p.Abbreviation,
		Name:         p.Name,
		City:         p.City,
		Country:      p.Country,
		Alias:        p.Alias,
		Regions:      p.Regions,
		Coordinates:  p.Coordinates,
		Province:     p.Province,
		Timezone:     p.Timezone,
		Unlocs:       p.Unlocs,
		Code:         p.Code,
	}
	if port.Alias == nil {
		port.Alias = []string{}
	}
	if port.Regions == nil {
		port.Regions = []string{}
	}
	if port.Unlocs == nil {
		port.Unlocs = []string{}
	}
	if port.Coordinates == nil {
		port.Coordinates = []float64{}
	}
	return port
}

// ImportResponse represents HTTP response for the port import call
type ImportResponse struct {
	NumberInserted  int32 `json:"number_inserted"`
	NumberUpdated   int32 `json:"number_updated"`
	EncounterErrors bool  `json:"encounter_errors"`
}
