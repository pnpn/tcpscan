package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

const (
	defaultAddress     string = "scanme.nmap.org"
	defaultStartPort   int    = 80
	defaultEndPort     int    = 80
	defaultConcurrency int    = 50
)

type Config struct {
	address     string
	startPort   int
	endPort     int
	concurrency int
}

func parseConfig() Config {
	address := flag.String("addresss", defaultAddress, "")
	port := flag.Int("port", defaultEndPort, "")
	concurrency := flag.Int("c", defaultConcurrency, "")
	flag.Parse()
	return Config{
		address:     *address,
		startPort:   *port,
		endPort:     *port,
		concurrency: *concurrency,
	}
}

func constructAddress(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}

func main() {
	config := parseConfig()
	_, err := net.Dial("tcp", constructAddress(config.address, config.startPort))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d open.\n", config.startPort)
}
