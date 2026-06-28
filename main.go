package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("--- Fenrir: Chaos Engineering Tool (Go) ---")
	fmt.Println("Simulating system stress...")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			fmt.Printf("[%s] Fenrir is hungry: consuming 100MB RAM (simulated)\n", time.Now().Format("15:04:05"))
			time.Sleep(2 * time.Second)
		}
	}()

	<-stop
	fmt.Println("\nFenrir is satiated. Shutting down.")
}
