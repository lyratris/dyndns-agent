package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

var (
	Data *Config
)

type Config struct {
	ID       string `ini:"apiID"`
	Key      string `ini:"apiKey"`
	Endpoint string `ini:"endpoint"`
	Interval int    `ini:"interval"`
	Protocol string `ini:"protocol"`
}

func Load() error {

	// Read environment variables
	envAPIID := os.Getenv("LTRS_DDNS_API_ID")
	envAPIKey := os.Getenv("LTRS_DDNS_API_KEY")
	envEndpoint := os.Getenv("LTRS_DDNS_ENDPOINT")
	envInterval := os.Getenv("LTRS_DDNS_INTERVAL")
	envProtocol := os.Getenv("LTRS_DDNS_PROTOCOL")

	if (envAPIID != "") || (envAPIKey != "") || (envEndpoint != "") || (envInterval != "") || (envProtocol != "") { // Use env variables if set

		// Make sure all required environment variables are defined
		if (envAPIID == "") || (envAPIKey == "") || (envEndpoint == "") {
			return fmt.Errorf("not all required environment variables are defined. Please make sure LTRS_DDNS_API_ID, LTRS_DDNS_API_KEY & LTRS_DDNS_ENDPOINTS are set")
		}

		// Set variabes
		var Data *Config
		Data.ID = envAPIID
		Data.Key = envAPIKey
		Data.Endpoint = envEndpoint
		Data.Interval, _ = strconv.Atoi(envInterval)
		Data.Protocol = envProtocol

	} else {

		// Read arguments
		argConfigPath := flag.String("config", "", "Path to config file")
		flag.Parse()

		// Define config path
		path := *argConfigPath
		if path == "" {
			path = "/etc/dyndns-agent/config.ini"
		}

		// Load the configuration file
		cfg, err := ini.Load(path)
		if err != nil {
			return fmt.Errorf("failed to load .ini file: %s", err)
		}

		// Unmarshal the .ini file into the Config struct
		var confData Config
		err = cfg.MapTo(&confData)
		if err != nil {
			return fmt.Errorf("failed to unmarshal .ini file: %s", err)
		}

		// Write to global var
		Data = &confData

	}

	// Set default variables for unset (or out of range) settings
	if Data.Interval < 120 || Data.Interval > 86400 {
		Data.Interval = 600
	}
	if !(Data.Protocol == "ipv4" || Data.Protocol == "ipv6" || Data.Protocol == "any") {
		Data.Protocol = "any"
	}

	return nil
}
