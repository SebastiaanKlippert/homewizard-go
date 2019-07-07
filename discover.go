package homewizard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type DiscoverResponse struct {
	IP     string `json:"ip"`
	Status string `json:"status"`
}

func DiscoverIP() (*net.IP, error) {

	resp, err := http.Get("http://gateway.homewizard.nl/discovery.php")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check if OK
	// TODO Better error handling
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("received HTTP statuscode %d, message: %s", resp.StatusCode, body)
	}

	// Unmarshal
	dr := new(DiscoverResponse)
	err = json.Unmarshal(body, dr)
	if err != nil {
		return nil, err
	}

	if dr.Status != "ok" {
		return nil, fmt.Errorf("received HomeWizard status %q", dr.Status)
	}

	pip := net.ParseIP(dr.IP)
	return &pip, nil
}
