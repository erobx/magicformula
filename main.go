package main

import (
	"fmt"

	"github.com/erobx/magicformula/parser"
	"github.com/erobx/magicformula/ui"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	ftp(err)

	app := ui.App{}
	app.Run("Magic Formula")

	//parser := parser.NewParser()

	//companies := parser.GetCompanies()
	//parser.Store(companies)
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
