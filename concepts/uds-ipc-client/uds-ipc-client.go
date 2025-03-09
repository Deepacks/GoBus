package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("unix", "/tmp/uds-ipc-client.sock")
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello Modbus Server"))
	if err != nil {
		log.Fatal("Write error:", err)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal("Read error:", err)
	}

	log.Printf("Received response: %s\n", string(buffer[:n]))
}