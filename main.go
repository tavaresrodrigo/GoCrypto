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

func main() {

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
	println(resp)

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
