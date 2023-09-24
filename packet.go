package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Define a struct to hold the decoded values
type DecodedPacket struct {
	Short1       uint16
	TwelveChars  string
	SingleByte   uint8
	EightChars   string
	Short2       uint16
	FifteenChars string
	Long         uint32
}

func decodePacket(packet []byte) (*DecodedPacket, error) {
	if len(packet) != 44 {
		return nil, fmt.Errorf("Invalid packet size, expected 44 bytes")
	}

	decoded := &DecodedPacket{}

	// Use binary.Read to decode binary data into the struct fields
	binary.Read(bytes.NewReader(packet[0:2]), binary.BigEndian, &decoded.Short1)
	decoded.TwelveChars = string(packet[2:14])
	decoded.SingleByte = packet[14]
	decoded.EightChars = string(packet[15:23])
	binary.Read(bytes.NewReader(packet[23:25]), binary.BigEndian, &decoded.Short2)
	decoded.FifteenChars = string(packet[25:40])
	binary.Read(bytes.NewReader(packet[40:44]), binary.BigEndian, &decoded.Long)

	return decoded, nil
}

func main() {
	// Sample input packet
	packet := []byte{0x04, 0xD2, 0x6B, 0x65, 0x65, 0x70, 0x64, 0x65, 0x63, 0x6F, 0x64, 0x69, 0x6E, 0x67, 0x38, 0x64, 0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70, 0x03, 0x15, 0x63, 0x6F, 0x6E, 0x67, 0x72, 0x61, 0x74, 0x75, 0x6C, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x73, 0x07, 0x5B, 0xCD, 0x15}

	// Decode the packet
	decoded, err := decodePacket(packet)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the decoded struct
	fmt.Printf("Decoded struct: %+v\n", decoded)
}