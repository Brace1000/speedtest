# 📡 Go Internet Speed Test Monitor

A lightweight Go application that **automates internet speed tests** and logs results to a CSV file. Ideal for developers, network admins, or anyone looking to monitor long-term internet performance.

---

## 🚀 Features

- 🕒 Runs a speed test every 10 seconds
- 📁 Saves results to `speedtest_results.csv`
- 🖥️ Prints real-time metrics to the terminal
- ✅ Graceful shutdown with `Ctrl+C`
- 📍 Logs: Timestamp, Download (Mbps), Upload (Mbps), Latency (ms), Server Name, Server Location

---
## ⚙️ Setup and Installation

### 1. Clone the Repository

```bash
git clone https://github.com/Brace1000/speedtest.git
cd speedtest
```

## 🛠️ Requirements

- Go 1.18+
- [`speedtest-go`](https://github.com/showwin/speedtest-go) library

Install the dependency:
```bash
go get github.com/showwin/speedtest-go/speedtest

/speedtest-monitor
├── main.go                # Entry point - sets up scheduler and handles shutdown
└── speed/
    └── performance.go     # Contains the speed test logic

```

 ## How It Works

    Fetches the best speed test server.

    Runs download, upload, and latency tests.

    Logs the results to both the console and a CSV file.

    Runs continuously every 10 seconds until interrupted.

## 💻 Sample Output

[2025-05-21 12:14:48]
Server: Nairobi (Kenya)
Download: 0.01 Mbps
Upload: 0.00 Mbps
Latency: 238.00 ms
--------------------------------------------------


## 🧹 Graceful Shutdown

When you press Ctrl+C, the app will:

    Stop the ticker

    Flush and close the CSV file

    Exit cleanly

 ##   📦 Possible Enhancements

    Export results to InfluxDB or Prometheus

    Visualize metrics with Grafana

    Alert system for slow connections

    Use cron for scheduling instead of Go ticker

  ##  Developed by Brian Oiko


---

Let me know if you want a version with badges, screenshots, or Docker instructions!
