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

	//The program first opens a text file containing all ISINs currently set up as valid instrument in the L&S investment universe (only equities, ex 'International equities'. After this, it retrieves current information about Open, High, Low, Close on yahoo.finance and stores the data into database

	//1. opens current investment universe data
	investmentUniverse := retrieveUniverse()
	fmt.Println(investmentUniverse)
	//2. open yahoo finance and crawl latest data
	crawlURI("BLA")
}

func retrieveUniverse() []string {

	//This function reads in the 'universe.txt' and returns a string array

	var universe []string //array which will be returned and passed into crawlURI()

	f, err := os.Open("universe.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() //closes the file after retrieveUniverse() returns

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		isin := scanner.Text()
		universe = append(universe, isin)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return universe

}

func crawlURI(isin string) {
	//This function gets passed ISIN, calls yahoo finance API and retrieves daily stock data

	fmt.Println(isin)
	// Variables
	yahooFinanceURI := "https://query1.finance.yahoo.com/v8/finance/chart/DAI.DE?symbol=DAI.DE?L&interval=1d"

	fmt.Println("Start crawling URI", yahooFinanceURI, "...")

	res, err := http.Get(yahooFinanceURI)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	//https: //www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
	//1. isins in yahoo symbols umwandeln,
	//2. yahoo daten ziehen..

}
