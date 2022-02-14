package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
)

type Response struct {
	Status string
	Code   string
	Total  int
	Data   []Coin
}

type Coin struct {
	Symbol  string
	CmcRank float64 `json:"cmc_rank"`
	Quote   Usd     `json:"quote"`
}

type Usd struct {
	Usd Price `json:"USD"`
}
type Price struct {
	Price float64 `json:"price"`
}

func captureHeader() string {
	cmk := os.Getenv("ccVar")
	return cmk
}

var coinOnDashboard = make([]map[string]string, 0)

func main() {

	greetUsers()

	// Api authentication.
	cmkValue := captureHeader()
	client := &http.Client{}
	serverUrl := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
	req, err := http.NewRequest("GET", serverUrl, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "50")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", cmkValue)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response Body")
		os.Exit(1)
	}

	data := Response{}
	json.Unmarshal(respBody, &data)
	fmt.Println(reflect.TypeOf(data))

}

// Function to add coins to the dashboard
func addCoin(coinSymbol string, coinAmount float64) {

	var userData = make(map[string]string)
	userData["coinSymbol"] = coinSymbol
	userData["Amount"] = strconv.FormatFloat(coinAmount, 'f', 2, 64)

	for {
		fmt.Print(`Inform the COIN symbol you want to add, ex: "BTC" for Bitcoin`)
		coinOnDashboard = append(coinOnDashboard, userData)
		fmt.Printf("Coin added to the map, there are all your coins being monitored: %v\n", coinOnDashboard)
	}

}

// Function to great users and show the main dashboard

func greetUsers() {
	fmt.Printf("Welcome to CommandCoin! Monitore your crypto assests from your terminal")

}

// TODO Create loop to iterate data.
