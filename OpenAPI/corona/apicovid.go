package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

		url := "https://covid-19-fastest-update.p.rapidapi.com/live/country/Indonesia/status/confirmed?date=2020-05-01T00%253A00%253A00Z"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "covid-19-fastest-update.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "<put your rapidapi key here>")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}