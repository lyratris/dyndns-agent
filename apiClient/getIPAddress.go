package apiClient

import (
	"encoding/json"
	"fmt"

	"forge.lyratris.com/lyratris-ltd/dyndns-agent/config"
)

func GetIPAddress() (string, string, error) {

	// Defaults
	ipv4 := "0.0.0.0"
	ipv6 := "::"

	if config.Data.Protocol == "ipv4" || config.Data.Protocol == "any" {

		// Get IPv4 Address
		ipv4ReqConfig := RequestConfig{
			Method: "GET",
			URL:    "https://api.lyratris.com/v1/other/ipaddress",
			IPType: "ipv4",
		}
		ipv4ReqResp, err := MakeRequest(ipv4ReqConfig)
		if err != nil {
			return "", "", err
		}

		// Parse reply
		var parsedV4Data struct {
			Address string `json:"address"`
		}
		err = json.Unmarshal([]byte(ipv4ReqResp), &parsedV4Data)
		if err != nil {
			return "", "", fmt.Errorf("error parsing API reply: %s", err)
		}

		ipv4 = parsedV4Data.Address

	}

	if config.Data.Protocol == "ipv6" || config.Data.Protocol == "any" {
		// Get IPv6 Address
		ipv6ReqConfig := RequestConfig{
			Method: "GET",
			URL:    "https://api.lyratris.com/v1/other/ipaddress",
			IPType: "ipv6",
		}
		ipv6ReqResp, err := MakeRequest(ipv6ReqConfig)
		if err != nil {
			return "", "", err
		}

		// Parse reply
		var parsedV6Data struct {
			Address string `json:"address"`
		}
		err = json.Unmarshal([]byte(ipv6ReqResp), &parsedV6Data)
		if err != nil {
			return "", "", fmt.Errorf("error parsing API reply: %s", err)
		}

		ipv6 = parsedV6Data.Address
	}

	return ipv4, ipv6, nil

}
