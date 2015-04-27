package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"os"
)

func main() {

	pcapFile := os.Args[1]
	if handle, err := pcap.OpenOffline(pcapFile); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			printAll(packet)
		}
	}
}

func printAll(pkt gopacket.Packet) {
	fmt.Println(pkt.Dump())
}
