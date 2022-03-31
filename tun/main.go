package main

import (
	"encoding/binary"
	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"
	"golang.org/x/net/icmp"
	"log"
)

func main() {
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Interface Name: %s\n", ifce.Name())

	buf := make([]byte, 2000)
	for {
		n, err := ifce.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		packet := buf[:n]
		src := waterutil.IPv4Source(packet)
		dest := waterutil.IPv4Destination(packet)
		log.Printf("Packet Received. len:%d src:%s dest:%s, payload:%v", len(packet), src.String(), dest.String(), packet)

		payload := waterutil.IPv4Payload(packet)

		typ := payload[0]

		id := int64(binary.LittleEndian.Uint16(payload[4:6]))

		seq := int64(binary.LittleEndian.Uint16(payload[6:8]))

		msg, err := icmp.ParseMessage(1, packet)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("icmp type:%v, cdoe:%v", msg.Type, msg.Code)

		log.Printf("type:%v, id:%v, seq:%v", typ, id, seq)
	}
}
