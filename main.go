package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"
)

type Coinsdata struct {
	Data []Coins `json:"data"`
	list int
}

type Coins struct {
	Name     string
	Cmc_rank int
	Quote    Quotes `json:"quote"`
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
	client := &http.Client{}

	serverUrl := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
	req, err := http.NewRequest("GET", serverUrl, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	
	buildRequest(req)

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

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintln(w, "Ranking\tName\tPrice\t")

	for _, data := range coinsData.Data {
		fmt.Fprintf(w, "%d\t%s\t%f\t\n", data.Cmc_rank, data.Name, data.Quote.USD.Price)
	}
}

func captureHeader() string {
	cmk := os.Getenv("CMMKVALUE")
	return cmk
}

func buildRequest(req *http.Request) {
	// Api authentication
	cmkValue := captureHeader()

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", cmkValue)

	// Getting the data passed from the cli parameters
	list := flag.Int("list", 20, "the number of coins to be listed")
	flag.Parse()

	// Converting from *int to string in order to build query parameter
	test := strconv.Itoa(*list)

	q := req.URL.Query()
	q.Add("limit", test)
	req.URL.RawQuery = q.Encode()
}
