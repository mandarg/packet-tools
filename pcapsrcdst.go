package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"os"
)

func main() {
	pcapFile := os.Args[1]
	if handle, err := pcap.OpenOffline(pcapFile); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		for pkt := range packetSource.Packets() {
			if ip4 := pkt.Layer(layers.LayerTypeIPv4); ip4 != nil {
				netFlow := pkt.NetworkLayer().NetworkFlow()
				src, dst := netFlow.Endpoints()
				fmt.Println(src)
				fmt.Println(dst)
			}
		}

	}
}
