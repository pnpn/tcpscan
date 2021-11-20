package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	defaultAddress     string = "scanme.nmap.org"
	defaultStartPort   string = "80"
	defaultEndPort     string = "80"
	defaultConcurrency int    = 50
)

type Config struct {
	address     string
	startPort   int
	endPort     int
	concurrency int
}

func parsePortString(ports string) (int, int) {
	var start, end string
	if strings.Contains(ports, "-") {
		portSplit := strings.Split(ports, "-")
		start = portSplit[0]
		end = portSplit[1]
	} else {
		start = ports
		end = ports
	}
	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)
	return startInt, endInt
}

func parseConfig() Config {
	address := flag.String("address", defaultAddress, "")
	port := flag.String("port", defaultEndPort, "")
	concurrency := flag.Int("c", defaultConcurrency, "")
	flag.Parse()
	startPort, endPort := parsePortString(*port)
	return Config{
		address:     *address,
		startPort:   startPort,
		endPort:     endPort,
		concurrency: *concurrency,
	}
}

func constructAddress(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}

func worker(address string, ports chan int, res chan int) {
	for port := range ports {
		d := net.Dialer{Timeout: 5 * time.Second}
		conn, err := d.Dial("tcp", constructAddress(address, port))
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
		go worker(config.address, portsChan, scanRes)
	}
	go func() {
		for port := config.startPort; port <= config.endPort; port++ {
			portsChan <- port
		}
	}()
	for i := config.startPort; i <= config.endPort; i++ {
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
