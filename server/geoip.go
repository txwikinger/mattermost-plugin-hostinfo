package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

type HostInfo struct {
	IP           string `json:"ip"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	RegionCode   string `json:"region_code"`
	RegionName   string `json:"region_name"`
	City         string `json:"city"`
	ZipCode      string `json:"zip_code"`
	TimeZone     string `json:"time_zone"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	MetroCode    int32   `json:"metro_code"`
}

type Consumer struct {
	serviceAddress string
	serviceEndpoint string
}

func GeoIP (host string) (hostInfo *HostInfo, err error) {
	c := Consumer{serviceAddress: "https://freegeoip.app", serviceEndpoint: "/json"}

	return c.GeoIP(host)
}

func (c Consumer) GeoIP(host string) (hostInfo *HostInfo, err error) {

	// ToDo: Refactor console printouts into log entries
	fmt.Println(c)

	url := c.serviceAddress + c.serviceEndpoint + "/" + host

	fmt. Println(url)
	
	// ToDo: Refactor - error handling
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// ToDo: log error
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	// ToDo: Refactor - error handling
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// ToDo: log error
		return nil, err
	}
	defer res.Body.Close()

	// ToDo: Refactor - error handling
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// ToDo: log error
		return nil, err
	}

	var hostinfo HostInfo

	bodyString := string(body)
	fmt.Println(bodyString)

	json.Unmarshal(body, &hostinfo)

	fmt.Println(hostinfo)

	// fmt.Println(res)

	return &hostinfo, nil
}

func (p *Plugin) renderHostinfoMessage(hostinfo *HostInfo, host string) *model.Post {

	message := "#### Host information for " + host + "\n"

	//message = message + "|:-----------|:-----------:|"
	message = message + "**Country Code:** " + hostinfo.CountryCode + "\n"
	message = message + "**Country Name:** " + hostinfo.CountryName + "\n"
	message = message + "**Region Code:** " + hostinfo.RegionCode + "\n"
	message = message + "**Region Name:** " + hostinfo.RegionName + "\n"
	message = message + "**City:** " + hostinfo.City + "\n"
	message = message + "**Zip Code:** " + hostinfo.ZipCode + "\n"
	message = message + "**Time Zone:** " + hostinfo.TimeZone + "\n"
	message = message + "**Latitude:** " + fmt.Sprintf("%f",hostinfo.Latitude) + "\n"
	message = message + "**Longitude:** " + fmt.Sprintf("%f",hostinfo.Longitude) + "\n"
	message = message + "**Metro Code:** " + fmt.Sprintf("%d",hostinfo.MetroCode) + "\n"

	post := &model.Post{
		Message: message,
		UserId:  p.botUserID,
	}

	return post

}

//func (p *Plugin) showHostinfoMessage(teamName string, args *model.CommandArgs, configMessage ConfigMessage) error {
func (p *Plugin) showHostinfoMessage(args *model.CommandArgs) error {
	host := strings.Fields(args.Command)[2]
	hostinfo, _ := GeoIP(host)
	post := p.renderHostinfoMessage(hostinfo, host)
	post.ChannelId = args.ChannelId
	_ = p.API.SendEphemeralPost(args.UserId, post)

	return nil
}
