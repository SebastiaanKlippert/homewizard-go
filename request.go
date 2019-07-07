package homewizard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httputil"
)

func (hw *HomeWizard) Do(in HWInput, out HWOutput) error {
	routeinfo := in.Route()

	// We have only implemented Get methods at this point
	if routeinfo.Method != "GET" {
		return fmt.Errorf("only GET methods are supported at this time")
	}
	if hw.IP.IsUnspecified() {
		return fmt.Errorf("IP address not set")
	}

	// Build URL
	uri := fmt.Sprintf("http://%s/%s%s", hw.IP.String(), hw.Password, routeinfo.Route)

	if hw.Verbose {
		log.Println("GET " + uri)
	}

	// Execute
	resp, err := hw.HTTPClient.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if hw.Verbose {
		b, _ := httputil.DumpResponse(resp, true)
		log.Println(string(b))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check if OK
	// TODO Better error handling
	if resp.StatusCode != 200 {
		return fmt.Errorf("received HTTP statuscode %d, message: %s", resp.StatusCode, body)
	}

	// Unmarshal into output
	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	if out.GetStatus() != "ok" {
		return fmt.Errorf("received HomeWizard status %q", out.GetStatus())
	}

	return nil
}
