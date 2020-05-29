package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "log"
import "net/http"
import "time"
import "strconv"


type Response struct {
	Confirmed struct{
		Value int `json:"value"`
		Detail string `json:"detail"`
	} `json:"confirmed"`

	Recovered struct{
		Value int `json:"value"`
		Detail string `json:"detail"`
	}  `json:"recovered"`

	Deaths Detail `json:"deaths"`

}

type Detail struct{
	Value int `json:"value"`
	Detail string `json:"detail"`
}


/*Name : httpRequest J*/
/*Detail : to execute  http request and get the body response*/
func httpRequest(url string) []byte {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }
    return body
}

/*Name : getAPIResponse Json object*/
/*Detail : to convert the http request to json object*/
func getAPIResponse(url string) Response{

	body := httpRequest(url)

	var response Response
	jsonErr := json.Unmarshal([]byte(body), &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return response
}

func main() {

	url := "https://covid19.mathdro.id/api/countries/IDN"

	response := getAPIResponse(url)

	fmt.Println("Covid-19 Data in Indonesia : ")
	fmt.Println("Number of confirmed 	: "+strconv.Itoa(response.Confirmed.Value))
	fmt.Println("Number of recovered 	: "+strconv.Itoa(response.Recovered.Value))
	fmt.Println("Number of Deaths 	: "+strconv.Itoa(response.Deaths.Value))

}