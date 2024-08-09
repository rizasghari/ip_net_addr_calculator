package ip

import (
	"net"
)

type IPV4 struct {
	Addr    net.IP
	NetMask net.IPMask
	NetAddr string
}

func (ip *IPV4) CalculateNetAddr() {
	if ip.Addr == nil {
		panic("Invalid IP or mask")
	}
	network := ip.Addr.Mask(ip.NetMask)
	ip.NetAddr = network.String()
}