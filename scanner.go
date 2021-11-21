package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func constructAddress(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}

func randomTimeToSleep() time.Duration {
	rand.Seed(time.Now().UnixNano())
	duration := rand.Float32() * 1000
	return time.Duration(duration) * time.Millisecond
}

func sleeper(config Config) {
	if config.randomWait {
		duration := randomTimeToSleep()
		fmt.Println(duration)
		time.Sleep(duration)
	} else {
		fmt.Println(config.waitTime)
		time.Sleep(time.Duration(config.waitTime) * time.Second)
	}

}

func worker(config Config, ports chan int, res chan int) {
	sleeper(config)
	for port := range ports {
		conn, err := net.DialTimeout("tcp", constructAddress(config.address, port), config.timeout*time.Second)
		if err != nil {
			res <- 0
			continue
		}
		conn.Close()
		res <- port
	}
}
