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

type Consumer struct {
	serviceAddress string
	serviceEndpoint string
}

func GeoIP (host string) *HostInfo {
	c := Consumer{serviceAddress: "https://freegeoip.app", serviceEndpoint: "/json"}

	return c.GeoIP(host)
}

func (c Consumer) GeoIP(host string) *HostInfo {

	// ToDo: Refactor console printouts into log entries
	fmt.Println(c)

	url := c.serviceAddress + c.serviceEndpoint + "/" + host

	fmt. Println(url)
	
	// ToDo: Refactor - error handling
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// ToDo: log error
		return nil
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	// ToDo: Refactor - error handling
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// ToDo: log error
		return nil
	}
	defer res.Body.Close()

	// ToDo: Refactor - error handling
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// ToDo: log error
		return nil
	}

	var hostinfo HostInfo

	bodyString := string(body)
	fmt.Println(bodyString)

	json.Unmarshal(body, &hostinfo)

	fmt.Println(hostinfo)

	// fmt.Println(res)

	return &hostinfo
}
