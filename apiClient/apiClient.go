package apiClient

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"forge.lyratris.com/lyratris-ltd/dyndns-agent/config"
)

// RequestConfig holds the configuration options for the HTTP request
type RequestConfig struct {
	Method string
	URL    string
	Data   any    // JSON data for POST requests
	IPType string // "ipv4", "ipv6", or "any"
}

// MakeRequest sends an HTTP request based on the provided configuration
func MakeRequest(reqConfig RequestConfig) (string, error) {
	// Prepare JSON data for POST requests
	var bodyReader *strings.Reader
	if reqConfig.Method == "POST" {
		jsonData, err := json.Marshal(reqConfig.Data)
		if err != nil {
			return "", err
		}
		bodyReader = strings.NewReader(string(jsonData))
	}

	// Create HTTP request
	var req *http.Request
	var err error
	if reqConfig.Method == "POST" {
		req, err = http.NewRequest(reqConfig.Method, reqConfig.URL, bodyReader)
		if err != nil {
			return "", err
		}
	} else if reqConfig.Method == "GET" {
		req, err = http.NewRequest(reqConfig.Method, reqConfig.URL, nil)
		if err != nil {
			return "", err
		}
	}

	// Set headers for POST requests
	if reqConfig.Method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Set authentication headers
	req.Header.Set("X-API-ID", config.Data.ID)
	req.Header.Set("X-API-Key", config.Data.Key)
	req.Header.Set("User-Agent", "Lyratris DynDNS Agent/1.0")

	// Set the DialContext function based on IPType
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	if reqConfig.IPType == "ipv6" {
		dialer.FallbackDelay = -1
	}

	// Perform the request
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				host, port, err := net.SplitHostPort(addr)
				if err != nil {
					return nil, err
				}

				switch reqConfig.IPType {
				case "ipv4":
					return dialer.DialContext(ctx, "tcp4", net.JoinHostPort(host, port))
				case "ipv6":
					return dialer.DialContext(ctx, "tcp6", net.JoinHostPort(host, port))
				default:
					return dialer.DialContext(ctx, network, addr)
				}
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read resp body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body), nil
}
