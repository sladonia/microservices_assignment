package domains

import (
	"port_domain_service/src/portpb"
)

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

func PortDomainFromPBPort(p *portpb.Port) *Port {
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
	if p.Coordinates == nil {
		p.Coordinates = []float64{}
	}
	if p.Alias == nil {
		p.Alias = []string{}
	}
	if p.Regions == nil {
		p.Regions = []string{}
	}
	if p.Unlocs == nil {
		p.Unlocs = []string{}
	}
	return port
}
