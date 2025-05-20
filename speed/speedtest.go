package speed

import (
      "log"
       "fmt"
	"github.com/showwin/speedtest-go/speedtest"
	"time"
	"encoding/csv"
)



func PerformTest(serverList speedtest.Servers, writer *csv.Writer) {
    // Find the best server
    targets, err := serverList.FindServer([]int{})
    if err != nil || len(targets) == 0 {
        log.Printf("Error finding server: %v\n", err)
        return
    }

    bestServer := targets[0]

   
    //err = bestServer.PingTest()
    if err != nil {
        log.Printf("Latency test failed: %v\n", err)
        return
    }

    // Run tests
    
    err = bestServer.DownloadTest()
    if err != nil {
        log.Printf("Download test failed: %v\n", err)
        return
    }

   
    err = bestServer.UploadTest()
    if err != nil {
        log.Printf("Upload test failed: %v\n", err)
        return
    }

    // Get and convert results
    downloadSpeed := bestServer.DLSpeed / 1000000 // Convert to Mbps
    uploadSpeed := bestServer.ULSpeed / 1000000   // Convert to Mbps
    latency := float64(bestServer.Latency.Milliseconds())
    serverName := bestServer.Name
    serverLocation := bestServer.Country

    // Prepare data for CSV
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    record := []string{
        timestamp,
        fmt.Sprintf("%.2f", downloadSpeed),
        fmt.Sprintf("%.2f", uploadSpeed),
        fmt.Sprintf("%.2f", latency),
        serverName,
        serverLocation,
    }

    // Write to CSV
    if err := writer.Write(record); err != nil {
        log.Printf("Error writing to CSV: %v\n", err)
    }
    writer.Flush()

    // Print results to console
    fmt.Printf("\n[%s]\n", timestamp)
    fmt.Printf("Server: %s (%s)\n", serverName, serverLocation)
    fmt.Printf("Download: %.2f Mbps\n", downloadSpeed)
    fmt.Printf("Upload: %.2f Mbps\n", uploadSpeed)
    fmt.Printf("Latency: %.2f ms\n", latency)
    fmt.Println("--------------------------------------------------")
}