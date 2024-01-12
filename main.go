package main

import (
	"log"
	"time"

	"forge.lyratris.com/lyratris-ltd/dyndns-agent/apiClient"
	"forge.lyratris.com/lyratris-ltd/dyndns-agent/config"
)

var (
	Addrv4 string
	Addrv6 string
)

func main() {

	// Load config
	err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	// Create a ticker that ticks every X seconds
	ticker := time.NewTicker(time.Second * time.Duration(config.Data.Interval))

	// Initial update
	run()

	// Infinite loop
	for {
		select {
		case <-ticker.C:
			run()
		}
	}
}

func run() {

	currentAddrv4, currentAddrv6, err := apiClient.GetIPAddress()
	if err != nil {
		log.Println(err)
		return
	}

	if (Addrv4 != currentAddrv4 && (config.Data.Protocol == "ipv4" || config.Data.Protocol == "any")) || (Addrv6 != currentAddrv6 && (config.Data.Protocol == "ipv6" || config.Data.Protocol == "any")) {

		// Update endpoint
		data := apiClient.EndpointData{
			IPv4: currentAddrv4,
			IPv6: currentAddrv6,
		}
		err = apiClient.UpdateEndpoint(data)
		if err != nil {
			log.Println(err)
			return
		}

		// Update cached addresses
		Addrv4 = currentAddrv4
		Addrv6 = currentAddrv6

		log.Printf("DynDNS Endpoint successfully udated. New IPv4: %s, New IPv6: %s", currentAddrv4, currentAddrv6)
	}

}
