package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HostInfo struct {
	IP           string
	Country_code string
	Country_name string
	Region_code  string
	Region_name  string
	City         string
	Zip_code     string
	Time_zone    string
	Latitude     float64
	Longitude    float64
	Metro_code   int32
}

func GeoIP(host string) HostInfo {

	url := "https://freegeoip.app/json/" + host

	// ToDo: Refactor - error handling
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	// ToDo: Refactor - error handling
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	// ToDo: Refactor - error handling
	body, _ := ioutil.ReadAll(res.Body)

	var hostinfo HostInfo

	bodyString := string(body)
	fmt.Println(bodyString)

	json.Unmarshal(body, &hostinfo)

	fmt.Println(hostinfo)

	// fmt.Println(res)

	return hostinfo
}
