package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"time"
	"encoding/json"
	"strconv"
)


type Row struct{
    Description string `json:"description"`
    DisplaySymbol string `json:"displaySymbol"`
    Symbol string `json:"symbol"`
}

type Price struct{
    Close float64 `json:"c"`
    High float64 `json:"h"`
    Low float64 `json:"l"`
    Open float64 `json:"o"`
    PClose float64 `json:"pc"`
    TimeStemp float64 `json:"t"`

}

type Response []Row 


type StockData struct{
	Description string `json:"description"`
    DisplaySymbol string `json:"displaySymbol"`
    Symbol string `json:"symbol"`
    Price Price `json:"price"`
}

type StockDataArray struct{
	StockData []StockData `json:"stockdata"`
}




func httpRequest(url string,headers [][]string) []byte {

    spaceClient := http.Client{
        Timeout: time.Second * 10, // Maximum of 5 secs
    }

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }

    for i := range headers{
            //fmt.Println(headers[i][0])
            //fmt.Println(headers[i][1])
            req.Header.Add(headers[i][0],headers[i][1])
    }

    res, getErr := spaceClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
        fmt.Println(getErr)
    }

    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }
    //fmt.Printf("%s", body)
    return body
}


//type Responses []Response

func getAPIResponseStock(url string,headers [][]string) Response{

    body := httpRequest(url,headers)

    var response Response
    jsonErr := json.Unmarshal([]byte(body), &response)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

    return response
}

func getAPIResponsePrice(url string,headers [][]string) Price{

    body := httpRequest(url,headers)

    var price Price
    jsonErr := json.Unmarshal([]byte(body), &price)
    if jsonErr != nil {
        log.Fatal(jsonErr)
        fmt.Println(jsonErr)
    }

    return price
}

func main() {

	var headers [][]string
    header1 := []string{ "x-rapidapi-host", "finnhub-realtime-stock-price.p.rapidapi.com" }
    header2 := []string{ "x-rapidapi-key", "2d6348677bmsh5d48dfc1a115e5ep159c37jsn9a08525aa294" }
    headers = append(headers,header1)
    headers = append(headers,header2)


	url := "https://finnhub-realtime-stock-price.p.rapidapi.com/stock/symbol?exchange=JK"

	response := getAPIResponseStock(url,headers)


	var stockDataArray StockDataArray

	for i := range response{
        //str := "Stock["+strconv.Itoa(i)+"] : "+response[i].Symbol+" : "+response[i].Description
        //fmt.Println(str)

        url = "https://finnhub-realtime-stock-price.p.rapidapi.com/quote?symbol="+response[i].Symbol
        price := getAPIResponsePrice(url,headers)

        stockData := StockData{
        				Description:response[i].Description,
        				DisplaySymbol:response[i].DisplaySymbol,
        				Symbol:response[i].Symbol,
        				Price:price}

       	byteArray, err := json.Marshal(stockData)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Stock["+strconv.Itoa(i)+"] : "+string(byteArray))
		stockDataArray.StockData = append(stockDataArray.StockData,stockData)
    	//fmt.Println("Price : "+strconv.FormatFloat(price.Close,'f', -1, 64))

    	if i % 50 == 0 && i != 0 {
    			fmt.Println("Sleeping zzzzzzzz....")
    		    time.Sleep(40 * time.Second)
    	}
		
    }

    byteArray, err := json.Marshal(stockDataArray)
	if err != nil {
		fmt.Println(err)
	}


	dt := time.Now()
    strdt:=dt.Format("2006-01-02T15:04:05")


	fmt.Println(string(byteArray))
	err = ioutil.WriteFile(strdt+"-stockData.json", byteArray, 0644)
	if err != nil {
    	fmt.Println(err)
	}
}