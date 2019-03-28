package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func String(IPAddr) string {
	var stringified string
	
	for i, _ := range IPAddr {
		stringified += (string(IPAddr[i]) + ".")
	}
	
	return stringified
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	
	
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip.String())
	}
}
