package main

import (
	"encoding/csv"
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

type stock struct {
	name string
	price float32
	change float32
}

func getStock(name string) stock {
	apiKey := "KKKMBBZJ4T7XPE4W"
	sanitizeName := url.QueryEscape(name)
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s&datatype=json",
						sanitizeName, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error in http request\n")
	}

	defer resp.Body.Close()
	
	var rawInfo info
	err = json.NewDecoder(resp.Body).Decode(&rawInfo)
	if err != nil {
		log.Fatal(err)
	}
	
	nprice, err := strconv.ParseFloat(rawInfo.Q.Price, 32)
	nchange, err := strconv.ParseFloat(rawInfo.Q.Change, 32)

	res := stock{
				rawInfo.Q.Tik,
				float32(nprice),
				float32(nchange),
			}
	return res
}


func main() {
	if len(os.Args) >= 3 && os.Args[1] == "-p" {
		port := os.Args[2]
		val := getPort(port)
		fmt.Printf("%.2f %+.2f\n", val.price, val.change)
		return
	}
	for _, v := range os.Args[1:] {
		val := getStock(v);
		fmt.Printf("%s: %.2f %+.2f\n", val.name, val.price, val.change)
	}

}

func getPort(name string) stock {
	hdir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path := hdir + string(os.PathSeparator) + "." + name + ".port"
	portFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	portcsv := csv.NewReader(portFile)
	portRecords, err := portcsv.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var price, change float32 
//	var count int	//for percentages 
	for _, v := range portRecords {
		num, err := strconv.Atoi(v[1])
		if err != nil {
			log.Fatal(err)
		}

		s := getStock(v[0])
//		count += num	//if we want to calculate percentage
		price += s.price * float32(num)
		change += s.change * float32(num)
	}
		
	return stock{name, price, change}
}
