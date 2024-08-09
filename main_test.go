package main

import (
	"net"
	"strconv"
	"strings"
	"testing"

	"github.com/rizasghari/net_addr_calculator/ip"
)

func TestNetAddrCalculator(t *testing.T) {

	tests := []struct {
		ipAddr   string
		netMask  string
		expected string
	}{
		{"78.191.15.209", "255.0.0.0", "78.0.0.0"},
		{"92.100.0.189", "128.0.0.0", "0.0.0.0"},
	}

	for _, test := range tests {

		netMaskStrArr := strings.Split(test.netMask, ".")
		var netMaskByteArr []byte
		for i := 0; i < 4; i++ {
			intVal, err := strconv.Atoi(netMaskStrArr[i])
			if err != nil {
				t.Fatal("Invalid netmask")
				return
			}
			netMaskByteArr = append(netMaskByteArr, byte(intVal))
		}

		ip := ip.IPV4{
			Addr:    net.ParseIP(test.ipAddr),
			NetMask: net.IPv4Mask(netMaskByteArr[0], netMaskByteArr[1], netMaskByteArr[2], netMaskByteArr[3]),
		}
		ip.CalculateNetAddr()
		if ip.NetAddr != test.expected {
			t.Fatalf("expected: %s, got: %s", test.expected, ip.NetAddr)
		}
	}
}
