package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			utilization, err := cpu.Percent(time.Second, true)
			if err != nil {
				slog.Error("error reading cpu", "err", err)
				continue
			}
			slog.Info("current usage", "utilization", utilization, "num_cpu", len(utilization))
		}
	}()

	// Setup a cancel and shutdown channelschannel
	c := make(chan os.Signal, 1)
	shutdown := make(chan bool)
	// Capture the interupt signal
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		stop(shutdown)
	}()

	// Wait for the shut down signal
	<-shutdown
}

func stop(shutdown chan bool) {
	// Cleanup services
	fmt.Printf("\nCleanup all connected services here...\n")

	// Send shutdown signal
	shutdown <- true
}
