package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

func captureHeader() string {
	cmk := os.Getenv("ccVar")
	return cmk
}

func main() {
	fetchData()
}

func fetchData() {

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

	fmt.Println(respBody)
	fmt.Printf("respBody variable type: %s\n\n", reflect.TypeOf(respBody))

}
