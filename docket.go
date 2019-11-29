package main

import (
	"net"
	"fmt"
	"bytes"
	"encoding/binary"
)

type headers struct {
	stream rune
	data []byte 
}

const socketLocation = "/var/run/docker.sock"

func read_header(head []byte, bytesRead int32) *headers {
	fmt.Printf("% x\n", head[:8])

	if head[0] == 1 {
		fmt.Printf("%d bytes from stdout\n", bytesRead)
	} else {
		fmt.Printf("%d bytes from stderr\n", bytesRead)
	}

	if head[1] != 0 || head[2] != 0 || head [3] != 0 {
		panic("not correct")
	}

	var bufLength int32
	buf := bytes.NewReader(head[4:8])
	err := binary.Read(buf, binary.BigEndian, &bufLength)

	header := headers{
		stream: 1,
		data: head[8:8+bufLength],
	}

	if err != nil {
		panic("guess I'll just die")
	}

	if (bytesRead != 8 + bufLength) {
		panic("Mismatch")
	}

	println(string(header.data))

	return &header
}

func main() {
	socket, err := net.Dial("unix", socketLocation) 	

	if err != nil {
		panic("fuck")
	}

	_, err = socket.Write([]byte("POST /containers/ee/attach?stream=1&stdout=1 HTTP/1.1\nHost:\n\n"))

	if err != nil {
		panic("fuck")
	}

	buff := make([]byte, 4096)
	_, _ = socket.Read(buff)

	for i := 0; i < 3; i++ {
		n, _ := socket.Read(buff)
		_ = read_header(buff, int32(n))
	}
}

