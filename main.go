package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


type Coinsdata struct {
	Data []Coins `json:"data"`
}

type Coins struct {
	Name string 
	Cmc_rank int 
	Quote Quotes `json:"quote"`
}

type Quotes struct {
	USD Prices `json:"USD"`
}

type Prices struct {
	Price float64 `json:"price"`
}

func main() {
	fetchData() 	
}

func fetchData() {
	// Api authentication
	cmkValue := captureHeader()
	client := &http.Client{}
	serverUrl := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
	req, err := http.NewRequest("GET", serverUrl, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", cmkValue)

	resp, errClient := client.Do(req)

	if errClient != nil {
		fmt.Println("Error sending request to server, ERROR: ", errClient)
		os.Exit(1)
	}

	respBody, errResponseBody := ioutil.ReadAll(resp.Body)

	if errResponseBody != nil {
		fmt.Println("Error reading response Body, ERROR: ", errResponseBody)
		os.Exit(1)
	}

	var coinsData Coinsdata
	err = json.Unmarshal([]byte(respBody), &coinsData) 
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	fmt.Println(coinsData)

	for i := 0; i < len(coinsData.Data); i++ {
        fmt.Println("Ranking: " , coinsData.Data[i].Cmc_rank)
		fmt.Println("Name: " , coinsData.Data[i].Name)
		fmt.Println("Price: " , coinsData.Data[i].Quote.USD.Price)

    }

}

func captureHeader() string {
	cmk := os.Getenv("CMMKVALUE")
	return cmk
}