package main

import (
	"fmt"
	//"io"
	"bufio"
	"log"
	"net/http"
	"os"
)

func main() {
	//1. opens current investment universe data
	retrieveUniverse()
	//2. open yahoo finance and crawl latest data
	crawlURI()
}

func retrieveUniverse() {

	var universe []string //array which will be returned and passed into crawlURI()

	f, err := os.Open("universe.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() //closes the file after retrieveUniverse() returns

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		isin := scanner.Text()
		universe = append(universe, isin)
	}

	fmt.Println(universe)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func crawlURI() {
	//This function gets passed ISINs and  calls yahoo finance API and retrieves daily stock data

	// Variables
	yahooFinanceURI := "https://query1.finance.yahoo.com/v8/finance/chart/DAI.DE?symbol=DAI.DE?L&interval=1d"

	fmt.Println("Start crawling URI", yahooFinanceURI, "...")

	res, err := http.Get(yahooFinanceURI)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}
