# Packet_Decoding-GoLang
In this code:

1. We define a DecodedPacket struct to hold the decoded values.
2. The decodePacket function takes the input packet as a byte slice, validates its size, and uses binary.Read to decode the binary data into the struct fields according to the provided decoding map.
3. In the main function, we provide a sample input packet and call the decodePacket function to decode it. The decoded struct is then printed.
