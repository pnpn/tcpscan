package main

import (
	"fmt"
	"net"
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
