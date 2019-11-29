package main

import (
	"net"
)

const socketLocation = "/var/run/docker.sock"

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
	for i := 0; i < 3; i++ {
		n, _ := socket.Read(buff)
		println(string(buff[0:n]))
	}
}
