package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func captureHeader() string {
	cmk := os.Getenv("CMMKVALUE")
	return cmk
}

type coinInfo struct {
	name 		string
	coinRanking float64
}

func main() {
	fetchData()
}

func fetchData() {
	var fetchdCoins[]coinInfo
	// Api authentication .
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

	var jsonData interface{}
	err = json.Unmarshal([]byte(respBody), &jsonData) // here!
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	allCoins := jsonData.(map[string]interface{})
	allCoinsData := allCoins["data"]
	allCoinsDataList := allCoinsData.([]interface{})

	for _, coins := range allCoinsDataList {
		coinMap := coins.(map[string]interface{})
		coin := coinInfo{name: coinMap["name"].(string), coinRanking: coinMap["cmc_rank"].(float64)}
		fetchdCoins = append(fetchdCoins, coin)
	}
	fmt.Println(fetchdCoins)
}
