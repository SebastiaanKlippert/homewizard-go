package homewizard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path/filepath"
)

type ConfigFile struct {
	IP       string
	Password string
	Verbose  bool
}

func NewHomeWizard(ip string, password string) *HomeWizard {
	return &HomeWizard{
		IP:         net.ParseIP(ip),
		Password:   password,
		HTTPClient: http.DefaultClient,
	}
}

func NewHomeWizardFromConfig(cfgFile string) (*HomeWizard, error) {
	config, err := ReadConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	if config.IP != "" {
		return NewHomeWizard(config.IP, config.Password), nil
	}
	hw, err := NewHomeWizardFromDiscovery(config.Password, config.Verbose)
	if err != nil {
		return nil, fmt.Errorf("no IP set in %s and discovery failed with error: %s", cfgFile, err)
	}
	return hw, nil
}

func NewHomeWizardFromDiscovery(password string, verbose bool) (*HomeWizard, error) {
	ip, err := DiscoverIP()
	if err != nil {
		return nil, err
	}
	if verbose {
		log.Printf("IP discovery OK, using IP %s\n", ip.String())
	}
	return &HomeWizard{
		Name:       "",
		IP:         *ip,
		Password:   password,
		Verbose:    verbose,
		HTTPClient: http.DefaultClient,
	}, nil
}

func ReadConfig(cfgFile string) (*ConfigFile, error) {
	buf, err := ioutil.ReadFile(filepath.Clean(cfgFile))
	if err != nil {
		return nil, err
	}
	config := new(ConfigFile)
	err = json.Unmarshal(buf, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
