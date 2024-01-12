package apiClient

import (
	"encoding/json"
	"fmt"

	"forge.lyratris.com/lyratris-ltd/dyndns-agent/config"
)

type EndpointData struct {
	IPv4 string `json:"IPv4"`
	IPv6 string `json:"IPv6"`
}

func UpdateEndpoint(data EndpointData) error {

	updateReqConfig := RequestConfig{
		Method: "POST",
		URL:    "https://api.lyratris.com/v1/dns/dynUpdate/" + config.Data.Endpoint,
		Data:   data,
		IPType: "any",
	}

	updateReqResp, err := MakeRequest(updateReqConfig)
	if err != nil {
		return err
	}

	// Parse reply
	var parsedResp struct {
		IPv4 string `json:"IPv4"`
		IPv6 string `json:"IPv6"`
	}
	err = json.Unmarshal([]byte(updateReqResp), &parsedResp)
	if err != nil {
		return fmt.Errorf("error parsing API reply: %s \n%s", err, updateReqResp)
	}

	if data.IPv4 != parsedResp.IPv4 || data.IPv6 != parsedResp.IPv6 {
		return fmt.Errorf("error during endpoint update, data missmatch")
	}

	return nil

}
