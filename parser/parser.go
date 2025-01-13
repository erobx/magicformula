package parser

import "golang.org/x/net/html"

type Company struct {
	Name              string
	Ticker            string
	MarketCap         string // $ Millions
	PriceFrom         string
	RecentQuarterData string
}

type Parser struct {
	doc *html.Node
}

func NewParser() Parser {
	var p Parser

	return p
}

func (p Parser) GetCompanies() []Company {

}
