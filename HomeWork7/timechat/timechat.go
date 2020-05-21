package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	fmt.Println("Starting server")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println("connected: ", conn.RemoteAddr().String())
		go handleConnRead(conn)
		go handleConn(conn)

	}
}
func handleConnRead(c net.Conn) {
	for {
		message, _ := bufio.NewReader(c).ReadString('\n')
		strMessage := strings.Replace(message, "\n", "", 2)
		strMessage = strings.Replace(strMessage, "\r", "", 2)

		fmt.Println("Received:", strMessage, " message")
		if strMessage == "exit" {
			fmt.Print("Stopping server")
			os.Exit(0)
		}
	}

}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
