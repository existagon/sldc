package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No SLD(s) specified")
		os.Exit(1)
	}

	sld := os.Args[1]

	checkSLD(sld)
}

func checkSLD(sld string) {
	tlds := getTLDs()

	var wg sync.WaitGroup

	for _, tld := range tlds {
		domain := strings.Join([]string{sld, ".", tld}, "")
		wg.Add(1)
		go checkDomain(domain, &wg)
	}

	wg.Wait()
}

func checkDomain(domain string, wg *sync.WaitGroup) {
	defer wg.Done()

	ips, err := net.LookupIP(domain)
	if err == nil {
		// TODO: dynamic padding
		fmt.Printf("%-20s | %v\n", domain, ips)
	}
}
