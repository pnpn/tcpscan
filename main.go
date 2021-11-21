package main

import (
	"fmt"
	"sort"
)

func main() {
	config := parseConfig()
	portsChan := make(chan int, config.concurrency)
	scanRes := make(chan int)
	var resSlice []int
	for i := 0; i <= cap(portsChan); i++ {
		go worker(config, portsChan, scanRes)
	}
	go func() {
		for _, port := range config.ports {
			portsChan <- port
		}
	}()
	for i := 0; i <= len(config.ports)-1; i++ {
		port := <-scanRes
		if port != 0 {
			resSlice = append(resSlice, port)
		}

	}
	close(portsChan)
	close(scanRes)
	sort.Ints(resSlice)
	for _, openPort := range resSlice {
		fmt.Printf("%v open.\n", openPort)
	}
}
