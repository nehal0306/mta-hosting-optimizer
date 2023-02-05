// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// )

// // IpConfig is the struct to hold IP configuration data
// type IpConfig struct {
// 	Hostname  string `json:"hostname"`
// 	IPAddress string `json:"ip_address"`
// 	Active    bool   `json:"active"`
// }

// // IpConfigs is a slice of IpConfig
// type IpConfigs []IpConfig

// var configs IpConfigs

// func main() {
// 	http.HandleFunc("/hosts", getHostsHandler)

// 	// Load IpConfig data from mock service
// 	err := loadIpConfigData()
// 	if err != nil {
// 		log.Fatalf("Failed to load IP config data: %v", err)
// 	}

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8081"
// 	}
// 	log.Printf("Starting server on port %s", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }

// func getHostsHandler(w http.ResponseWriter, r *http.Request) {
// 	xStr := r.URL.Query().Get("x")
// 	fmt.Println("in handler", xStr)
// 	if xStr == "" {
// 		xStr = os.Getenv("DEFAULT_X")
// 		fmt.Println("is null - ", xStr)
// 		if xStr == "" {
// 			xStr = "1"
// 		}
// 	}
// 	x, err := strconv.Atoi(xStr)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Invalid value for x: %v", err), http.StatusBadRequest)
// 		return
// 	}

// 	hosts := map[string]int{}
// 	for _, config := range configs {
// 		if config.Active {
// 			hosts[config.Hostname]++
// 		}
// 	}

// 	result := []string{}
// 	for host, count := range hosts {
// 		if count <= x {
// 			result = append(result, host)
// 		}
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(result)
// }

// func loadIpConfigData() error {
// 	// Load data from mock service
// 	// ...

// 	return nil
// }
