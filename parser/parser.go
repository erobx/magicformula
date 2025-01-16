package parser

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

type Parser struct {
	doc *html.Node
}

type Company struct {
	Name              string
	Ticker            string
	MarketCap         string // $ Millions
	PriceFrom         string
	RecentQuarterData string
}

type RequestBody struct {
	MinimumMarketCap string
	Select30         string
	Stocks           string
}

func NewRequstBody(cap, select30, stocks string) RequestBody {
	return RequestBody{
		MinimumMarketCap: cap,
		Select30:         select30,
		Stocks:           stocks,
	}
}

func NewParser() Parser {
	var p Parser

	res := makeRequest()
	doc, err := html.Parse(res.Body)
	ftp(err)

	p.doc = doc
	return p
}

func (p Parser) GetCompanies() []Company {
	companies := make([]*Company, 0)
	companies = processCompanies(p.doc, companies)
	companies = companies[len(companies)-50:] // trim other tr elements that result in whitespace

	// return as values not pointers
	newComps := make([]Company, 0)
	for _, c := range companies {
		newComp := *c
		newComps = append(newComps, newComp)
	}

	return newComps
}

func (p Parser) Store(companies []Company) {
	now := time.Now()
	y, m, d := now.Date()
	title := fmt.Sprintf("data/%d-%d-%d.csv", m, d, y)

	headers := []string{"Company Name", "Ticker", "Market Cap ($ Millions)", "Price From", "Most Recent Quarter Data"}
	records := [][]string{
		headers,
	}

	for _, c := range companies {
		record := []string{c.Name, c.Ticker, c.MarketCap, c.PriceFrom, c.RecentQuarterData}
		records = append(records, record)
	}

	f, err := os.OpenFile(title, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(records)
	if err := w.Error(); err != nil {
		log.Fatalln("Error writing to csv:", err)
	}

	fmt.Println("Saved stocks to csv")
}

func processCompanies(n *html.Node, companies []*Company) []*Company {
	if n.Type == html.ElementNode && n.Data == "tr" {
		comp := processNode(n)
		companies = append(companies, comp)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		companies = processCompanies(c, companies)
	}
	return companies
}

func processNode(n *html.Node) *Company {
	comp := &Company{}
	it := 0
	switch n.Data {
	case "tr":
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			for t := c.FirstChild; t != nil; t = t.NextSibling {
				if t.Type == html.TextNode {
					selectCompanyValue(it, t, comp)
				}
				it++
			}
		}
	}
	return comp
}

func selectCompanyValue(it int, t *html.Node, comp *Company) {
	switch it {
	case 0:
		comp.Name = t.Data
	case 1:
		comp.Ticker = t.Data
	case 2:
		comp.MarketCap = t.Data
	case 3:
		comp.PriceFrom = t.Data
	case 4:
		comp.RecentQuarterData = t.Data
	}
}

func makeRequest() *http.Response {
	fmt.Println("Requesting stocks...")

	url := "https://magicformulainvesting.com/Screening/StockScreening"
	cookie := os.Getenv("COOKIE")

	client := &http.Client{}

	reqBody := NewRequstBody("100", "false", "Get+Stocks") // always gets 50 despite bool val
	out, err := json.Marshal(reqBody)
	ftp(err)

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))
	ftp(err)
	req.Header.Set("Cookie", cookie)

	res, err := client.Do(req)
	ftp(err)
	return res
}

func ftp(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
