// package tcp.
package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// listenAddr binds every interface so players on other machines can join.
const listenAddr = ":4242"

// TCP listens on listenAddr and serves each client in its own goroutine.
// It only returns when the listener cannot be created.
func TCP() error {
	//nolint:gosec // public on purpose, see listenAddr.
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", listenAddr, err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Printf("close listener: %v", err)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error acception connection:", err)
			continue
		}
		fmt.Println("Server → Client: OK hello proto=1")

		go handleConnection(conn)
	}
}

func parsecommand(message string) (string, string) {
	command, after, _ := strings.Cut(strings.TrimSpace(message), " ")
	return strings.ToUpper(command), after
}

func writeMessage(title string, message string, conn net.Conn) {
	response := fmt.Sprintf(title, message)
	_, err := conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server writer error: %v", err)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("close connection: %v", err)
		}
	}()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}
		command, data := parsecommand(message)
		writeMessage("ACK: %s\n", command, conn)
		if command == "QUIT" {
			writeMessage("%s\n", "BYE:", conn)
			fmt.Println("QUIT THE GAME")
			return
		}
		fmt.Println(command, data)
	}
}
