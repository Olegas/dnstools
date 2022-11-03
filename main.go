package main

import (
	"flag"
	"fmt"

	"github.com/Olegas/dnstools/pkg/config"
	"github.com/Olegas/dnstools/pkg/dns"
	"github.com/Olegas/dnstools/pkg/ip"
)

var dryRun = flag.Bool("d", false, "Dry run")

func main() {
	flag.Parse()
	ip := ip.GetCurrentIp()
	fmt.Printf("Current IP: %s", ip)
	if !*dryRun {
		res := dns.UpdateRecord(ip, config.GetZoneID())
		fmt.Printf("Result: %t", res)
	}
}
