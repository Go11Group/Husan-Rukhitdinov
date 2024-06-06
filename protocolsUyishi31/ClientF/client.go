package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server.
	conn, err := net.Dial("tcp", "localhost:8070")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	go sender(conn)

	listener(conn)
}

func sender(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("your message: ", input)
	}
}

func listener(conn net.Conn) {
	for {
		byt := make([]byte, 1000)
		n, err := conn.Read(byt)
		if err != nil {
			fmt.Println("Error while listening message: ", err.Error())
		}
		fmt.Println(string(byt[:n]))
	}
}
