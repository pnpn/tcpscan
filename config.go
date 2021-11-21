package main

import (
	"flag"
	"strconv"
	"strings"
)

const (
	defaultAddress        string = "scanme.nmap.org"
	defaultConcurrency    int    = 50
	defaultMaxConcurrency bool   = false
	defaultTimeout        int    = 5
	defaultPort           string = "80"
)

type Config struct {
	address        string
	concurrency    int
	maxConcurrency bool
	timeout        time.Duration
	ports          []int
}

func parsePortString(ports string) []int {
	var portsSplit []string
	var portsIntList []int
	if strings.Contains(ports, ",") {
		portsSplit = strings.Split(ports, ",")
	}
	for _, portsSection := range portsSplit {
		if strings.Contains(portsSection, "-") {
			start, end := strings.Split(portsSection, "-")[0], strings.Split(portsSection, "-")[1]
			startInt, _ := strconv.Atoi(start)
			endInt, _ := strconv.Atoi(end)
			for i := startInt; i <= endInt; i++ {
				portsIntList = append(portsIntList, i)
			}
		} else {
			portInt, _ := strconv.Atoi(portsSection)
			portsIntList = append(portsIntList, portInt)

		}

	}
	sort.Ints(portsIntList)
	return portsIntList
}

func parseConfig() Config {
	address := flag.String("address", defaultAddress, "")
	port := flag.String("port", defaultPort, "")
	concurrency := flag.Int("c", defaultConcurrency, "")
	maxConcurrency := flag.Bool("max-c", defaultMaxConcurrency, "")
	timeout := flag.Int("t", defaultTimeout, "")
	flag.Parse()
	ports := parsePortString(*port)
	if *maxConcurrency {
		*concurrency = len(ports)
	}
	return Config{
		address:     *address,
		concurrency: *concurrency,
		timeout:     time.Duration(*timeout) * time.Second,
		ports:       ports,
	}
}
