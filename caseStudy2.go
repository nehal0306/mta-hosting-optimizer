package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type IpConfig struct {
	IP     string `json:"ip"`
	Host   string `json:"hostname"`
	Active bool   `json:"active"`
}

func main() {
	http.HandleFunc("/optimizeServer", optimizeServer)
	err := http.ListenAndServe(":4001", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getIPConfig() []IpConfig {
	//this can be treated as a mock service.
	return []IpConfig{
		{IP: "127.0.0.1", Host: "mta-prod-1", Active: true},
		{IP: "127.0.0.2", Host: "mta-prod-1", Active: false},
		{IP: "127.0.0.3", Host: "mta-prod-2", Active: true},
		{IP: "127.0.0.4", Host: "mta-prod-2", Active: true},
		{IP: "127.0.0.5", Host: "mta-prod-2", Active: false},
		{IP: "127.0.0.6", Host: "mta-prod-3", Active: false},
	}
}

func optimizeServer(w http.ResponseWriter, r *http.Request) {
	//Taking X fron Environment variables
	xStr := os.Getenv("X")

	var x int
	if xStr == "" {
		fmt.Println("set to 1")
		x = 1
	} else {
		var err error
		x, err = strconv.Atoi(xStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	ipConfig := getIPConfig()
	//creating a hashmap to keep track of
	hostCount := make(map[string]int)
	for _, config := range ipConfig {
		if _, ok := hostCount[config.Host]; !ok {
			hostCount[config.Host] = 0
		}
		if config.Active {
			hostCount[config.Host]++
		}
	}

	var hosts []string
	for host, count := range hostCount {
		if count <= x {
			hosts = append(hosts, host)
		}
	}

	hostsJSON, err := json.Marshal(hosts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(hostsJSON)
}
