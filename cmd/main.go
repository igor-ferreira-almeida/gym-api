package main

import (
	"fmt"
	"github.com/gym-api/cmd/server"
	"time"
)

func main() {
	start := time.Now()
	LinearSearch([]string{"1", "2", "3", "4", "5", "6"}, "7")
	timeElapsed := time.Since(start)
	fmt.Printf("linear search time elapsed: %s", timeElapsed)
	// go bigqueue.Receive()
	server.NewServer()
}

func LinearSearch(array []string, targetValue string) int {
	for guess := 0; guess < len(array); guess++ {
		if array[guess] == targetValue {
			return guess
		}
	}
	return -1
}
