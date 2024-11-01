package main

import (
	"net"
	"testing"
)

func TestIpCov(t *testing.T) {
	ip := net.ParseIP("127.0.0.1")
	ipn := ip.To4()
	t.Log(ipn)
	t.Log(ipn[0], ipn[1], ipn[2], ipn[3])
	ipNum := uint32(ipn[0])<<24 + uint32(ipn[1])<<16 + uint32(ipn[2])<<8 + uint32(ipn[3])
	t.Log(ipNum)
}
