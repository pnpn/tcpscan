package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func constructAddress(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}

func worker(address string, timeout time.Duration, ports chan int, res chan int) {
	for port := range ports {
		conn, err := net.DialTimeout("tcp", constructAddress(address, port), timeout*time.Second)
		if err != nil {
			res <- 0
			continue
		}
		conn.Close()
		res <- port
	}
}

func main() {
	config := parseConfig()
	portsChan := make(chan int, config.concurrency)
	scanRes := make(chan int)
	var resSlice []int
	for i := 0; i <= cap(portsChan); i++ {
		go worker(config.address, config.timeout, portsChan, scanRes)
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
