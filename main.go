package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
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

func main() {
	config := parseConfig()
	for i := config.startPort; i <= config.endPort; i++ {
		conn, err := net.Dial("tcp", constructAddress(config.address, i))
		if err != nil {
			continue
		} else {
			fmt.Printf("%d open.\n", i)
		}
		conn.Close()
	}
}
