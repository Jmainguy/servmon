package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func checkHTTPService(servHTTP servHTTP, retryMap map[string]int) (servHTTP, map[string]int) {
	resp, err := http.Get(servHTTP.URL)
	if err != nil {
		log.Println(err)
		servHTTP.Status = "down"
		retryMap[servHTTP.URL]++
		return servHTTP, retryMap
	}
	if servHTTP.Statuscode == resp.StatusCode {
		servHTTP.Status = "up"
		retryMap[servHTTP.URL] = 0
	} else {
		servHTTP.Status = "down"
		retryMap[servHTTP.URL]++

	}
	// Check error from resp.Body.Close()
	if cerr := resp.Body.Close(); cerr != nil {
		log.Println("error closing response body:", cerr)
	}
	return servHTTP, retryMap
}

func checkTCPService(servTCP servTCP, retryMap map[string]int) (servTCP, map[string]int) {
	timeout := time.Second
	for pk, port := range servTCP.Ports {
		serviceName := fmt.Sprintf("%s-%s-%s", servTCP.Hostname, port.Network, port.Port)
		conn, err := net.DialTimeout(port.Network, net.JoinHostPort(servTCP.Hostname, port.Port), timeout)
		if err != nil {
			servTCP.Ports[pk].Status = "down"
			retryMap[serviceName]++
		} else {
			if conn != nil {
				servTCP.Ports[pk].Status = "up"
				retryMap[serviceName] = 0
				// Check error from conn.Close()
				if cerr := conn.Close(); cerr != nil {
					log.Println("error closing connection:", cerr)
				}
			}
		}
	}
	return servTCP, retryMap
}
