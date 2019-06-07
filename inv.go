package main

import (
//	"encoding/csv"
	"encoding/json"
	"fmt"
//	"io"
	"log"
	"net/url"
	"net/http"
	"os"
	"strconv"
)

type info struct {
	Q struct {
		Tik string `json:"01. symbol"`
		Price string `json:"05. price"`
		Change string `json:"09. change"`
	} `json:"Global Quote"`
}

type record struct {
	name string
	price float32
	change float32
}

func main() {
	apiKey := "KKKMBBZJ4T7XPE4W"
	ticker := ""
	if len(os.Args) == 2 {
		ticker = os.Args[1]
	}
	
	sanitizeTicker := url.QueryEscape(ticker)
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s&datatype=json",
						sanitizeTicker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error in http request\n")
	}

	defer resp.Body.Close()
	
	var sInfo info
	err = json.NewDecoder(resp.Body).Decode(&sInfo)
	if err != nil {
		log.Fatal(err)
	}
	
	nprice, err := strconv.ParseFloat(sInfo.Q.Price, 32)
	nchange, err := strconv.ParseFloat(sInfo.Q.Change, 32)

	stock := record{
				sInfo.Q.Tik,
				float32(nprice),
				float32(nchange),
			}
				
	fmt.Printf("%s: %.2f %+.2f\n", stock.name, stock.price, stock.change)
}
