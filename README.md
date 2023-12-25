### SERVER
 - The 'handleConnection function' is responsible for handling each incoming client connection.
 It reads data from the client and responds with a simple message.

 - The 'main' function sets up a TCP server that listens to port 8080. It continuously accepts
   incoming connections and handles each connection in a seperate goroutine.

### CLIENT

- The 'main' function in the client connects to the server using 'net.Dial'. It sends a message to
  the server and reads the response.
