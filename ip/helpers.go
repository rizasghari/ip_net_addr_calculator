package ip

import (
	"fmt"
	"strconv"
	"strings"
)

func ipToBinary(ip string) []string {
	octets := strings.Split(ip, ".")
	var binary []string
	for _, octet := range octets {
		num, err := strconv.Atoi(octet)
		catchError(err)
		binaryOctet := fmt.Sprintf("%08b", num)
		binary = append(binary, binaryOctet)
	}

	return binary
}

func binaryToIP(binary []string) string {
	var ip []string

	for _, binaryOctet := range binary {
		num, err := strconv.ParseUint(binaryOctet, 2, 8)
		if err != nil {
			fmt.Println("Invalid binary IP")
			return ""
		}
		ip = append(ip, fmt.Sprintf("%d", num))
	}

	return strings.Join(ip, ".")
}

func catchError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetNetMaskBytes(netMask string) []byte {
	netMaskStrArr := strings.Split(netMask, ".")
	var netMaskByteArr []byte
	for _, netMask := range netMaskStrArr {
		intVal, err := strconv.Atoi(netMask)
		if err != nil {
			fmt.Println("Invalid netmask")
			return nil
		}
		netMaskByteArr = append(netMaskByteArr, byte(intVal))
	}
	return netMaskByteArr
}