package main

import (
	"bufio"
	"fmt"
	"net"
)

var mp = make(map[net.Conn]bool)

func main() {
	// Listen for incoming connections.
	listener, err := net.Listen("tcp", ":8070")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8070...")
	for {
		// Accept a connection from a client.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) // Handle each connection in a new goroutine.
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())
	mp[conn] = true
	reader := bufio.NewReader(conn)
	for {

		fmt.Println(mp)
		// Read data from the connection.
		message, err := reader.ReadString('\n')
		if err != nil {
			delete(mp, conn)
			fmt.Println("Error reading message:", err)
			return
		}

		// Echo the message back to the client..
		fmt.Print("Received message: ", string(message))
		message = message[:len(message)-1]

		// fmt.Println(ls[i])
		// _, err = conn.Write([]byte(strings.ToUpper(message) + " FROM SERVER!\n")) // Echo back in uppercase.
		// if err != nil {
		// 	fmt.Println("Error writing message:", err)
		// 	return
		// }
		for i := range mp {
			fmt.Println("sende_________")
			_, err := i.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("Error while sending message from server: ", err.Error())
			}
		}
	}

}
