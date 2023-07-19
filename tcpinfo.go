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
	handle, err := pcap.OpenOffline(pcapFile)

	if err != nil {
		panic(err)
	}

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			ip, _ := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
			fmt.Printf("%d → %d\n", ip.SrcIP, ip.DstIP)
			fmt.Printf("Sequence Number: %d\n", tcp.Seq)
			fmt.Printf("Window Size: %d\n\n", tcp.Window)
		}
	}
}
