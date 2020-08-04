package main

import (
	"testing"
)

// ToDo: Refactor: Better test (not requiring connection to Internet)
func TestGeoIP(t *testing.T) {

	host := "dns.google.com"

	result := GeoIP(host)

	if result.Country_code != "US" {
		t.Errorf("Resulting Country Code was %s instead of %s\n", result.Country_code, "US")
	}
	if result.Country_name != "United States" {
		t.Errorf("Resulting Country Code was %s instead of %s\n", result.Country_name, "United States")
	}
	/*
			if result.Region_code != "" {
				t.Errorf("Resulting Country Code was %s instead of %s\n", result.Region_code, "")
			}
			if result.Region_name != "" {
				t.Errorf("Resulting Country Code was %s instead of %s\n", result.Region_name, "")
			}
			if result.Zip_code != "" {
				t.Errorf("Resulting Country Code was %s instead of %s\n", result.Zip_code, "")
			}
			if result.Time_zone != "America/Chicago" {
				t.Errorf("Resulting Country Code was %s instead of %s\n", result.Country_code, "America/Chicago")
			}
			if !(result.Latitude == 37.751) {
				t.Errorf("Resulting Country Code was %f instead of %s\n", result.Latitude, "37.751")
			}
			if !(result.Longitude == -97.822) {
				t.Errorf("Resulting Country Code was %f instead of %s\n", result.Longitude, "-97.822")
			}
			if !(result.Metro_code == 0) {
				t.Errorf("Resulting Country Code was %d instead of %s\n", result.Metro_code, "0")
		       }
	*/

}
