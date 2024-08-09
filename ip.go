package main

import (
	"fmt"
	"strconv"
)

type IP struct {
	IPV4    string
	netMask string
	netAddr string
}

func (ip *IP) calculateNetAddr() {
	ipBinary := ipToBinary(ip.IPV4)
	maskBinary := ipToBinary(ip.netMask)
	if ipBinary == nil || maskBinary == nil {
		panic("Invalid IP or mask")
	}

	var networkBinary []string

	for i := 0; i < 4; i++ {
		ipOctet, _ := strconv.ParseUint(ipBinary[i], 2, 8)
		maskOctet, _ := strconv.ParseUint(maskBinary[i], 2, 8)
		networkOctet := ipOctet & maskOctet
		networkBinary = append(networkBinary, fmt.Sprintf("%08b", networkOctet))
	}

	ip.netAddr = binaryToIP(networkBinary)
}