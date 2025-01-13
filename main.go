package main

import (
	"fmt"

	"github.com/erobx/magicformula/parser"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	ftp(err)

	parser := parser.NewParser()

	companies := parser.GetCompanies()
	for _, comp := range companies {
		printCompany(comp)
	}
}

func printCompany(comp parser.Company) {
	fmt.Printf("%s, %s, %s\n", comp.Name, comp.Ticker, comp.MarketCap)
}

func ftp(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
