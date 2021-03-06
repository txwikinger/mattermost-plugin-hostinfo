package main

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/json/dns.google.com", googleMock)
 
	srv := httptest.NewServer(handler)
 
	return srv
}

func googleMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{
			"ip":"8.8.8.8",
			"country_code":"US",
			"country_name":"United States",
			"region_code":"IL",
			"region_name":"Illinois",
			"city":"Chicago",
			"zip_code":"12345",
			"time_zone":"America/Chicago",
			"latitude":37.751,
			"longitude":-97.822,
			"metro_code":0}`))
}

func TestGeoIP(t *testing.T) {

	srv := serverMock()
	defer srv.Close()
 
	assert := assert.New(t)

	c := Consumer{serviceAddress: srv.URL, serviceEndpoint: "/json"}

	host := "dns.google.com"

	result := c.GeoIP(host)

	assert.Equal("8.8.8.8", result.IP)
	assert.Equal("US", result.Country_code)
	assert.Equal("United States", result.Country_name)
	assert.Equal("IL", result.Region_code)
	assert.Equal("Illinois", result.Region_name)
	assert.Equal("Chicago", result.City)
	assert.Equal("12345", result.Zip_code)
	assert.Equal("America/Chicago", result.Time_zone)
	assert.Equal(37.751, result.Latitude)
	assert.Equal(-97.822, result.Longitude)
	assert.Equal(int32(0), result.Metro_code)
}
