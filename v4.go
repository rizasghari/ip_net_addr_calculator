package main

import (
	"fmt"
	"strconv"
)

type IPV4 struct {
	addr    string
	netMask string
	netAddr string
}

func (ip *IPV4) calculateNetAddr() {
	ipBinary := ipToBinary(ip.addr)
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