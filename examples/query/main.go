package main

import (
	"context"
	"fmt"
	"net"

	"github.com/takc923/mdns"
	"golang.org/x/net/ipv4"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", mdns.DefaultAddress)
	if err != nil {
		panic(err)
	}

	l, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}

	server, err := mdns.Server(ipv4.NewPacketConn(l), &mdns.Config{})
	if err != nil {
		panic(err)
	}
	answer, src, err := server.ReverseLookup(context.TODO(), []byte{192, 168, 1, 30})
	fmt.Println(answer)
	fmt.Println(src)
	fmt.Println(err)
}
