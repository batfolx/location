package location

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// struct that holds the details of the ip
type IPDetails struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	ORG         string  `json:"osg"`
	AS          string  `json:"as"`
}

type IPError struct {
	message string
}

func (e *IPError) Error() string {
	return (*e).message
}

func GetLocationDetails(ipAddress string) (IPDetails, error) {

	// get the endpoint
	ipError := IPError{message: "none"}
	endpoint := fmt.Sprintf("http://ip-api.com/json/%s", ipAddress)
	resp, err := http.Get(endpoint)
	if err != nil {
		ipError.message = err.Error()
		return IPDetails{}, &ipError
	}

	// make ip details
	details := IPDetails{}

	// convert data to bytes
	data, _ := ioutil.ReadAll(resp.Body)

	// unmarshal data into data structure
	if err := json.Unmarshal(data, &details); err != nil {
		ipError.message = err.Error()
		return IPDetails{}, &ipError
	}

	// return ip details
	return details, nil

}
