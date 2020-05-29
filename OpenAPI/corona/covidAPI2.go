package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "log"
import "net/http"
import "time"
import "strconv"
import "image/color"

import "gonum.org/v1/plot"
import "gonum.org/v1/plot/plotter"
import "gonum.org/v1/plot/vg"
import "gonum.org/v1/plot/vg/draw"

import "os/exec"




type Row struct{
	Country string `json:"Country"`
	CountryCode string `json:"CountryCode"`
	Province string `json:"Province"`
	City string `json:"City"`
	CityCode string `json:"CityCode"`
	Lat string `json:"Lat"`
	Lon string `json:"Lon"`
	Confirmed float64 `json:"Confirmed"`
	Active float64 `json:"Active"`
	Deaths float64 `json:"Deaths"`
	Recovered float64 `json:"Recovered"`
	Date string `json:"Date"`
}

type Response []Row 




/*Name : httpRequest J*/
/*Detail : to execute  http request and get the body response*/
func httpRequest(url string,headers [][]string) []byte {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for i := range headers{
			fmt.Println(headers[i][0])
			fmt.Println(headers[i][1])
			req.Header.Add(headers[i][0],headers[i][1])
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("%s", body)
    return body
}

/*Name : getAPIResponse Json object*/
/*Detail : to convert the http request to json object*/
func getAPIResponse(url string,headers [][]string) Response{

	body := httpRequest(url,headers)

	var response Response
	jsonErr := json.Unmarshal([]byte(body), &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return response
}



func generatePoints(n int,data Response,plotStr string) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
			pts[i].X =  float64(i)
		switch plotStr {
		case "Confirmed" :
			pts[i].Y = data[i].Confirmed
		case "Active" :
			pts[i].Y = data[i].Active
		case "Deaths" :
			pts[i].Y = data[i].Deaths
		case "Recovered" :
			pts[i].Y = data[i].Recovered

		}
	}
	return pts
}

func generatePointsLine(data Response,plotStr string) (*plotter.Line, *plotter.Scatter, error) {
	
	// Make a line plotter and set its style.
	lpLine, lpPoints, err := plotter.NewLinePoints(generatePoints(len(data),data,plotStr))
	if err != nil {
		panic(err)
	}
	
	lpLine.Color = color.RGBA{G: 255, A: 255}

	switch plotStr {
		case "Confirmed" :
			lpLine.Color = color.RGBA{B: 255, A: 255}
		case "Active" :
			lpLine.Color = color.RGBA{B: 255, A: 255}
		case "Deaths" :
			lpLine.Color = color.RGBA{R: 255, B: 128, A: 255}
		case "Recovered" :
			lpLine.Color = color.RGBA{G: 255, A: 255}
	}


	lpPoints.Shape = draw.PyramidGlyph{}
	lpPoints.Color = color.RGBA{R: 255, A: 255}

	return lpLine, lpPoints, err
}


func plotData (data Response){

	var lpLine *plotter.Line 
	var lpPoints *plotter.Scatter

	fmt.Println("Plotting Data.......")

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Covid-19 Data in Indonesia"
	p.X.Label.Text = "Days"
	p.Y.Label.Text = "Count"
	p.Legend.Top = true
	p.Legend.Left = true


	lpLine, lpPoints, err   = generatePointsLine(data,"Active")
	if err != nil {
		panic(err)
	}
	
	p.Add(lpLine)
	p.Legend.Add("Active", lpLine, lpPoints)


	lpLine, lpPoints, err   = generatePointsLine(data,"Recovered")
	if err != nil {
		panic(err)
	}
	
	p.Add(lpLine)
	p.Legend.Add("Recovered", lpLine, lpPoints)


	lpLine, lpPoints, err   = generatePointsLine(data,"Deaths")
	if err != nil {
		panic(err)
	}
	
	p.Add(lpLine)
	p.Legend.Add("Deaths", lpLine, lpPoints)




	// Save the plot to a PNG file.
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "Covid-19.png"); err != nil {
		panic(err)
	}

	out, err := exec.Command("/usr/bin/qlmanage", "-p", "Covid-19.png").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }

    out=out
    //output := string(out[:])
    //fmt.Println(output)
    fmt.Println("Command Successfully Executed")
}




func main() {

	var headers [][]string
	header1 := []string{ "x-rapidapi-host", "covid-19-fastest-update.p.rapidapi.com" }
	header2 := []string{ "x-rapidapi-key", "<put your rapidapi key here>" }
	//header3 := []string{ "useQueryString", "true"}


	url := "https://covid-19-fastest-update.p.rapidapi.com/live/country/Indonesia/status/confirmed?date=2020-03-01T00%253A00%253A00Z"
	headers = append(headers,header1)
	headers = append(headers,header2)
	//headers = append(headers,header3)

	response := getAPIResponse(url,headers)

	fmt.Println("Covid-19 Data in Indonesia : ")
	for i := range response{
		str := "Date   of confirmed["+strconv.Itoa(i)+"] 	: "+response[i].Date
		//fmt.Println(str)
		str =str
	}

	plotData(response)

}