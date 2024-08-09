package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/rizasghari/net_addr_calculator/ip"
)

func main() {
	var ipAddr string
	fmt.Println("Enter IPv4: ")
	fmt.Scanf("%s", &ipAddr)

	var netMask string
	fmt.Println("Enter netmask: ")
	fmt.Scanf("%s", &netMask)

	netMaskStrArr := strings.Split(netMask, ".")

	var netMaskByteArr []byte
	for i := 0; i < 4; i++ {
		intVal, err := strconv.Atoi(netMaskStrArr[i])
		if err != nil {
			fmt.Println("Invalid netmask")
			return
		}
		netMaskByteArr = append(netMaskByteArr, byte(intVal))
	}

	v4 := ip.IPV4{
		Addr:    net.ParseIP(ipAddr),
		NetMask: net.IPv4Mask(netMaskByteArr[0], netMaskByteArr[1], netMaskByteArr[2], netMaskByteArr[3]),
	}
	v4.CalculateNetAddr()
	fmt.Printf("%+v\n", v4)

	// var ip net.IP
	// var ipAddr string
	// fmt.Println("Enter IPv4: ")
	// fmt.Scanf("%s", &ipAddr)
	// ip = net.ParseIP(ipAddr)
	// if ip == nil {
	// 	fmt.Println("Invalid IP")
	// 	return
	// }
	// fmt.Printf("%+v\n", ip.String())

}
