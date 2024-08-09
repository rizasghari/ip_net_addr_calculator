package main

import (
	"fmt"
)

func main() {
	var ipAddr string
	fmt.Println("Enter IPv4: ")
	fmt.Scanf("%s", &ipAddr)

	var netMask string
	fmt.Println("Enter netmask: ")
	fmt.Scanf("%s", &netMask)

	ipv4 := IPV4{addr: ipAddr, netMask: netMask}
	ipv4.calculateNetAddr()
	fmt.Printf("%+v\n", ipv4)
}
