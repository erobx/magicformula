package main

import (
	"fmt"
	"os"

	"github.com/erobx/magicformula/parser"
	"github.com/erobx/magicformula/ui"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected scrape flag")
		os.Exit(1)
	}

	err := godotenv.Load()
	ftp(err)

	switch os.Args[1] {
	case "y":
		fmt.Println("Scrapping html...")
		runParser()
	case "n":
		fmt.Println("Starting app..")
		runApp()
	}
}

func runApp() {
	app := ui.NewApp()
	app.Run("Magic Formula")
}

func runParser() {
	parser := parser.NewParser()
	companies := parser.GetCompanies()
	parser.Store(companies)
}

func ftp(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
