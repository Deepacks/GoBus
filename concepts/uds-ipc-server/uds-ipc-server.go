package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("unix", "/tmp/uds-ipc-client.sock")

	if (err != nil) {
		log.Fatal("Listen error: ", err)
	}

	defer listener.Close()
	log.Println("Server is listening on /tmp/uds-ipc-client.sock...")

	for {
		conn, err := listener.Accept()

		if (err != nil) {
			log.Fatal("Listen error: ", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	
	if err != nil {
		log.Println("Read error:", err)
		return
	}

	log.Printf("Received: %s\n", string(buffer[:n]))

	_, err = conn.Write([]byte("ACK"))
	if err != nil {
		log.Println("Write error:", err)
		return
	}
}
