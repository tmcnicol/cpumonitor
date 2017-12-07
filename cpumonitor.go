package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {

	ticker := time.NewTicker(time.Second)
	fmt.Println(ticker)
	go func() {
		for t := range ticker.C {
			utilization, err := cpu.Percent(time.Second, false)
			if err != nil {
				fmt.Println("Something fucked up")
			}
			fmt.Println(t, utilization)

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
