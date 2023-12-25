package main 

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Handle the connection: read and write data
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return 
	} 

	fmt.Println("Received: %s/n", buffer)

	// respond to the client
	response := []byte("Hello from the server!")
	conn.Write(response)
}

func main() {
	// listen to port 8080
	listener, err := net.Listen("top", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return 
	}
	defer listener.Close()

	fmt.Println("Server listening on :8080")

	for {
		// Accept a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}