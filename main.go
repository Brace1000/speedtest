package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
	"speed/speed"
)

func main() {
	
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Create or open CSV file
	file, err := os.OpenFile("speedtest_results.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header if file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}

	if fileInfo.Size() == 0 {
		headers := []string{"Timestamp", "Download (Mbps)", "Upload (Mbps)", "Latency (ms)", "Server Name", "Server Location"}
		if err := writer.Write(headers); err != nil {
			log.Fatalf("Failed to write headers: %v", err)
		}
		writer.Flush()
	}

	// Set up graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	
	// Fetch user info (required first)
_, err = speedtest.FetchUserInfo()
if err != nil {
	log.Fatalf("Error fetching user info: %v", err)
}

// Get server list
serverList, err := speedtest.FetchServers()
if err != nil {
	log.Fatalf("Error fetching servers: %v", err)
}


	// For continuous testing
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	fmt.Println("Starting speed tests every 10 seconds. Press Ctrl+C to stop.")
	fmt.Println("Results will be saved to speedtest_results.csv")
	fmt.Println("--------------------------------------------------")

	// Perform initial test immediately
	speed.PerformTest(serverList, writer)

	// Main loop
	for {
		select {
		case <-ticker.C:
			speed.PerformTest(serverList, writer)
		case <-signalChan:
			fmt.Println("\nGracefully shutting down...")
			fmt.Println("Final results saved to speedtest_results.csv")
			return
		}
	}
}

