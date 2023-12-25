package main 

import (
	"fmt"
	"net"
)

func main() {
	// connect to the server
	conn, err := net.Dial("top", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return 
	}
	defer conn.Close()


	// Send data to the server
	message := []byte("Hello from the client!")
	conn.Write(message)

	// Receive and print the response
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Printf("Response from server: %s\n", buffer)
}