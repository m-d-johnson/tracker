package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/prototext"
	"log"
	"os"
)

import pb "tracker/proto"

func main() {
	titan := pb.Device{
		Name:        "Desktop Computer",
		Hostname:    "titan",
		Model:       "Custom",
		Description: "Mike's Desktop Computer",
		Interfaces: []*pb.Device_NetworkInterface{
			{
				Name:      "Wired Ethernet",
				Mac:       "112233445566",
				Ip:        "192.168.20.8",
				MediaType: "Copper",
				Speed:     "1000 Mbps",
			},
			{
				Name:      "Wireless Ethernet",
				Mac:       "AA2233445566",
				Ip:        "192.168.20.9",
				MediaType: "Radio",
				Speed:     "600 Mbps",
			}},
	}
	fmt.Print(titan.Interfaces[0])
	// Write the new address book back to disk.
	out, err := proto.Marshal(&titan)
	textProto, _ := prototext.Marshal(&titan)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile("data.bin.pb", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	if err := os.WriteFile("data.textproto", textProto, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
